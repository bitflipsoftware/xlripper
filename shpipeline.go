package xlripper

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/bitflip-software/xlsx/xmlprivate"
)

var rowRoutines = runtime.NumCPU()
var cellRoutines = runtime.NumCPU()

type topInfo struct {
	runes  []rune
	shared sharedStrings
}

type rowInfo struct {
	top          topInfo
	interationIX int
	rowLoc       tagLoc
}

type cellInfo struct {
	rowIX   int
	colIX   int
	cellLoc tagLoc
	rowInfo rowInfo
}

type cellParseResult struct {
	cellInfo
	value *string
}

type rowParseResult struct {
	rowInfo
	cells []cellParseResult
}

type rowChan chan rowParseResult
type cellChan chan cellParseResult

func parseRowAsync(r rowInfo, ch rowChan, wg *sync.WaitGroup) {
	ch <- parseRow(r)
	wg.Done()
}

func parseRow(r rowInfo) rowParseResult {
	ix := shSetLast(r.top.runes, r.rowLoc.open.last+1)
	e := shSetLast(r.top.runes, r.rowLoc.close.first-1)
	ch := make(cellChan, cellRoutines)
	sendWait := sync.WaitGroup{}
	receiveWait := sync.WaitGroup{}
	rpr := rowParseResult{}
	rpr.top.runes = r.top.runes
	rpr.rowLoc = r.rowLoc
	rpr.interationIX = r.interationIX
	rpr.cells = make([]cellParseResult, 0)
	receiveWait.Add(1)
	go receiveCellsAsync(ch, &rpr, &receiveWait)

cellLoop:
	for {
		openLoc := shFindFirstOccurenceOfElement(r.top.runes, ix, e, "c")

		if openLoc == badPair {
			break cellLoop
		}

		closeLoc := shTagCloseFind(r.top.runes, openLoc.last+1, e, "c")

		if closeLoc == badPair {
			break cellLoop
		}

		cellLoc := tagLoc{openLoc, closeLoc}
		c := cellInfo{}
		c.rowInfo = r
		c.rowIX = -1
		c.colIX = -1
		c.cellLoc = cellLoc

		sendWait.Add(1)
		go parseCellAsync(c, ch, &sendWait)
		ix = cellLoc.close.last + 1
	}

	sendWait.Wait()
	close(ch)
	receiveWait.Wait()
	return rpr
}

func receiveRowsAsync(ch rowChan, outSheet *Sheet, wg *sync.WaitGroup) {
	for r := range ch {
		for _, c := range r.cells {
			outSheet.add(c.rowIX, c.colIX, c.value)
		}
	}
	wg.Done()
}

func receiveCellsAsync(ch cellChan, outRowResult *rowParseResult, wg *sync.WaitGroup) {
	for c := range ch {
		if c.rowIX >= 0 && c.colIX >= 0 {
			outRowResult.cells = append(outRowResult.cells, c)
		}
	}
	wg.Done()
}

func parseCellAsync(c cellInfo, ch cellChan, wg *sync.WaitGroup) {
	ch <- parseCell(c)
	wg.Done()
}

func parseCell(c cellInfo) cellParseResult {
	str := string(c.rowInfo.top.runes[c.cellLoc.open.first : c.cellLoc.close.last+1])
	xmlC := xmlprivate.C{}
	err := xml.Unmarshal([]byte(str), &xmlC)

	if err != nil {
		// TODO - introduce pipeline cancellation on err
		fmt.Print(err.Error())
	}

	result := cellParseResult{}
	result.rowInfo = c.rowInfo
	result.cellLoc = c.cellLoc
	result.cellInfo = c

	rowIX, colIX := parseRowIndexCellIndex(xmlC.R)
	result.rowIX = rowIX
	result.colIX = colIX

	if xmlC.T == "s" {
		// should be a shared string
		if sharedIX, err := strconv.Atoi(xmlC.V); err == nil {
			shStr := c.rowInfo.top.shared.get(sharedIX)
			result.value = shStr
		}
	} else {
		result.value = &xmlC.V
	}

	return result
}

func parseRowIndexCellIndex(sheetCellReference string) (rowIX, colIX int) {
	letterBuf := bytes.Buffer{}
	numberBuf := bytes.Buffer{}

	for _, r := range strings.ToUpper(sheetCellReference) {
		if r >= 'A' && r <= 'Z' {
			letterBuf.WriteRune(r)
		} else if r >= '0' && r <= '9' {
			numberBuf.WriteRune(r)
		} else {
			return -1, -1
		}
	}

	rowIX, err := strconv.Atoi(numberBuf.String())
	rowIX--

	if err != nil {
		return -1, -1
	}

	colIX = lettersToNum(letterBuf.String())
	return rowIX, colIX
}

func letterToNum(r rune) int {
	if r < 'A' || r > 'Z' {
		return -1
	}

	i := int(r - 'A')

	if i < 0 || i > 25 {
		return -1
	}

	return i
}

func lettersToNum(str string) int {
	if len(str) == 1 {
		rs := []rune(str)
		if len(rs) == 1 {
			return letterToNum(rs[0])
		}
	} else if len(str) == 0 {
		return -1
	}

	nums := make([]int, 0, 2)
	for _, rn := range str {
		n := letterToNum(rn)
		if n < 0 || n > 25 {
			return -1
		}
		num := n
		nums = append(nums, num)
	}

	exp := len(nums) - 1
	sum := 0

	for _, n := range nums {
		if exp == 0 {
			sum += n
		} else {
			add := int(math.Pow(float64(26), float64(exp))) * (n + 1)
			//cur := add + n
			sum += add
			exp--
		}
	}

	return sum
}
