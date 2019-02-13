package xlripper

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var rowRoutines = runtime.NumCPU() + 1
var cellRoutines = maxi(runtime.NumCPU()/4, 2)

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
	rpr.cells = make([]cellParseResult, 0, 65)
	receiveWait.Add(1)
	go receiveCellsAsync(ch, &rpr, &receiveWait)

cellLoop:
	for {
		openLoc, isSelfClosing := shFindFirstOccurenceOfElement(r.top.runes, ix, e, "c")

		if openLoc == badPair {
			break cellLoop
		}

		closeLoc := badPair

		if !isSelfClosing {
			closeLoc, _ = shTagCloseFind(r.top.runes, openLoc.last+1, e, "c")
		} else {
			closeLoc = indexPair{openLoc.last, openLoc.last}
		}

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
		parseCellAsync(c, ch, &sendWait)
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
	runes := c.rowInfo.top.runes[c.cellLoc.open.first : c.cellLoc.close.last+1]
	//xmlC := xmlprivate.CellXML{}
	core := cellCoreFast{}
	err := core.parseXML(runes)
	//err := xml.Unmarshal([]byte(str), &xmlC)

	//if rand.Float64() < 0.00001 {
	//	f, err := os.OpenFile("/Users/mjb/Desktop/c_test.go", os.O_APPEND|os.O_WRONLY, 0600)
	//	if err != nil {
	//		panic(err)
	//	}
	//js, _ := json.Marshal(xmlC)
	//
	//	strBld := bytes.Buffer{}
	//	strBld.WriteRune('{')
	//	strBld.WriteRune('`')
	//	strBld.WriteString(str)
	//	strBld.WriteRune('`')
	//	strBld.WriteRune(',')
	//	strBld.WriteRune('`')
	//	strBld.Write(js)
	//	strBld.WriteRune('`')
	//	strBld.WriteRune(',')
	//	strBld.WriteRune('}')
	//	strBld.WriteRune(',')
	//	strBld.WriteString("\n")
	//
	//	defer f.Close()
	//
	//	if _, err = f.WriteString(strBld.String()); err != nil {
	//		panic(err)
	//	}
	//
	//	use(f)
	//
	//}

	if err != nil {
		// TODO - introduce pipeline cancellation on err
		fmt.Print(err.Error())
	}

	result := cellParseResult{}
	result.rowInfo = c.rowInfo
	result.cellLoc = c.cellLoc
	result.cellInfo = c

	rowIX, colIX := parseRowIndexCellIndex(core.cellReference())
	result.rowIX = rowIX
	result.colIX = colIX

	if core.typeInfo() == ctSharedString {
		// should be a shared string
		if sharedIX, err := strconv.Atoi(*core.value()); err == nil {
			shStr := c.rowInfo.top.shared.get(sharedIX)
			result.value = shStr
		}
	} else {
		result.value = core.value()
	}
	//else {
	//	if len(xmlC.V) > 0 {
	//		result.value = &xmlC.V
	//	} else if len(xmlC.InlineString.Str) > 0 {
	//		result.value = &xmlC.InlineString.Str
	//	} else {
	//		result.value = &emptyString
	//	}
	//
	//}

	return result
}

func parseRowIndexCellIndex(sheetCellReference string) (rowIX, colIX int) {
	//letterBuf := bytes.Buffer{}
	//numberBuf := bytes.Buffer{}

	//for _, r := range strings.ToUpper(sheetCellReference) {
	//	if r >= 'A' && r <= 'Z' {
	//		letterBuf.WriteRune(r)
	//	} else if r >= '0' && r <= '9' {
	//		numberBuf.WriteRune(r)
	//	} else {
	//		return -1, -1
	//	}
	//}

	up := strings.ToUpper(sheetCellReference)

	firstNum := -1
	for ix, r := range up {
		if r >= '0' && r <= '9' {
			firstNum = ix
			break
		}
	}

	if firstNum <= 0 {
		return -1, -1
	}

	letters := up[:firstNum]
	nums := up[firstNum:]

	rowIX, err := strconv.Atoi(nums)
	rowIX--

	if err != nil {
		return -1, -1
	}

	colIX = lettersToNum(letters)

	if colIX < 0 {
		return -1, -1
	}

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
