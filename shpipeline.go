package xlsx

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"sync"

	"github.com/bitflip-software/xlsx/xmlprivate"
)

type topInfo struct {
	runes  []rune
	shared sharedStrings
}

type rowInfo struct {
	top    topInfo
	rowIX  int
	rowLoc tagLoc
}

type cellInfo struct {
	cellIX  int
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
	ch := make(cellChan, 16)
	wg := sync.WaitGroup{}
	rpr := rowParseResult{}
	rpr.top.runes = r.top.runes
	rpr.rowLoc = r.rowLoc
	rpr.rowIX = r.rowIX
	rpr.cells = make([]cellParseResult, 0)
	go receiveCellsAsync(ch, &rpr)

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
		c.cellIX = -1
		c.cellLoc = cellLoc

		wg.Add(1)
		go parseCellAsync(c, ch, &wg)
		ix = cellLoc.close.last + 1
	}

	wg.Wait()
	close(ch)

	return rpr
}

func receiveRowsAsync(ch rowChan, outSheet *Sheet) {
	for r := range ch {
		for ix, c := range r.cells {
			outSheet.add(r.rowIX, ix, c.value)
		}
	}
}

func receiveCellsAsync(ch cellChan, outRowResult *rowParseResult) {
	for c := range ch {
		if c.cellIX > 0 {
			for i := len(outRowResult.cells); i < c.cellIX; i++ {
				outRowResult.cells = append(outRowResult.cells, cellParseResult{})
			}

			outRowResult.cells[c.cellIX] = c
		}
	}
}

func parseCellAsync(c cellInfo, ch cellChan, wg *sync.WaitGroup) {
	ch <- parseCell(c)
	wg.Done()
}
func parseCell(c cellInfo) cellParseResult {
	str := string(c.rowInfo.top.runes[c.cellLoc.open.first : c.cellLoc.close.last+1])
	fmt.Print(str)
	xmlC := xmlprivate.C{}
	err := xml.Unmarshal([]byte(str), &xmlC)

	if err != nil {
		//?
		fmt.Print(err.Error())
	}

	result := cellParseResult{}
	result.rowInfo = c.rowInfo
	result.cellLoc = c.cellLoc
	result.cellInfo = c

	result.cellIX = 2 // TODO parse the cell index

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

func parseRowIndexCellIndex(sheetCellReference string) (rowIX, cellIX int) {
	return -1, -1
}

func letterToNum(r rune) int {
	if r < 'A' || r > 'Z' {
		return -1
	}

	i := int('A' - r)
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
	for rn := range []rune(str) {
		num := letterToNum(rune(rn))
		nums = append(nums, num)
	}

	return -1
}
