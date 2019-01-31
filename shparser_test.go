package xlripper

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestUnitShParserBasics(t *testing.T) {
	tn := "TestShParserBasics"
	rawData := topen(Mac1621)
	zs, err := zopen(rawData)

	if err != nil {
		t.Errorf("received error from zopen: '%s'", err.Error())
		return
	}

	sh, err := shparse(zs, 0)

	if err != nil {
		t.Errorf("received error from shparse: '%s'", err.Error())
		return
	}

	sheetName := "^^^^ gs"
	if sh.Name != sheetName {
		t.Errorf(tfail(tn, "sh.Name", sh.Name, sheetName))
	}

	sheetIX := 0
	if sh.Name != sheetName {
		t.Errorf(tfail(tn, "sh.Index", itos(sh.Index), itos(sheetIX)))
	}

	got := itos(len(sh.Columns))
	want := "10"
	stmt := "itos(len(sh.Columns))"
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX := 0
	expectedCellCount := 4
	col := sh.Columns[colIX]
	colStmt := fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 1
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 2
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 3
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 4
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 5
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 6
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 7
	expectedCellCount = 4
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 8
	expectedCellCount = 0
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	colIX = 9
	expectedCellCount = 6
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	got = itos(len(col.Cells))
	want = itos(expectedCellCount)
	stmt = fmt.Sprintf("len(%s.Cells)", colStmt)
	if got != want {
		t.Error(tfail(tn, stmt, got, want))
		return
	}

	rowIX := 5
	colIX = 9
	expectedCellValue := "z"
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	cell := col.Cells[rowIX]
	cellStmt := fmt.Sprintf("Cells[%d]", rowIX)
	stmt = fmt.Sprintf("*(%s.%s)", colStmt, cellStmt)
	if cell == nil {
		t.Errorf("%s would result in a nil dereference, the value '%s' was expected instead", stmt, expectedCellValue)
	} else {
		got = *cell
		want = expectedCellValue
		if got != want {
			t.Error(tfail(tn, stmt, got, want))
		}
	}

	rowIX = 0
	colIX = 0
	expectedCellValue = "sfd"
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	cell = col.Cells[rowIX]
	cellStmt = fmt.Sprintf("Cells[%d]", rowIX)
	stmt = fmt.Sprintf("*(%s.%s)", colStmt, cellStmt)
	if cell == nil {
		t.Errorf("%s would result in a nil dereference, the value '%s' was expected instead", stmt, expectedCellValue)
	} else {
		got = *cell
		want = expectedCellValue
		if got != want {
			t.Error(tfail(tn, stmt, got, want))
		}
	}

	rowIX = 1
	colIX = 2
	expectedCellFloat := 57.37072
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	cell = col.Cells[rowIX]
	cellStmt = fmt.Sprintf("Cells[%d]", rowIX)
	stmt = fmt.Sprintf("*(%s.%s)", colStmt, cellStmt)
	if cell == nil {
		t.Errorf("%s would result in a nil dereference, the value '%s' was expected instead", stmt, expectedCellValue)
	} else {
		got = *cell
		gotF, err := strconv.ParseFloat(got, 64)

		if err != nil {
			t.Errorf("error trying to parse %s as a float: %s", got, err.Error())
		}

		want = fmt.Sprintf("%.14f", expectedCellFloat)
		if got != want {
			if math.Abs(gotF-expectedCellFloat) > epsilon {
				t.Error(tfail(tn, stmt, got, want))
			}
		}
	}

	rowIX = 2
	colIX = 3
	expectedCellValue = "sad;fghjk"
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	cell = col.Cells[rowIX]
	cellStmt = fmt.Sprintf("Cells[%d]", rowIX)
	stmt = fmt.Sprintf("*(%s.%s)", colStmt, cellStmt)
	if cell == nil {
		t.Errorf("%s would result in a nil dereference, the value '%s' was expected instead", stmt, expectedCellValue)
	} else {
		got = *cell
		want = expectedCellValue
		if got != want {
			t.Error(tfail(tn, stmt, got, want))
		}
	}

	rowIX = 3
	colIX = 7
	expectedCellFloat = 8.73734
	col = sh.Columns[colIX]
	colStmt = fmt.Sprintf("sh.Columns[%d]", colIX)
	cell = col.Cells[rowIX]
	cellStmt = fmt.Sprintf("Cells[%d]", rowIX)
	stmt = fmt.Sprintf("*(%s.%s)", colStmt, cellStmt)
	if cell == nil {
		t.Errorf("%s would result in a nil dereference, the value '%s' was expected instead", stmt, expectedCellValue)
	} else {
		got = *cell
		gotF, err := strconv.ParseFloat(got, 64)

		if err != nil {
			t.Errorf("error trying to parse %s as a float: %s", got, err.Error())
		}

		want = fmt.Sprintf("%.14f", expectedCellFloat)
		if got != want {
			if math.Abs(gotF-expectedCellFloat) > epsilon {
				t.Error(tfail(tn, stmt, got, want))
			}
		}
	}
}

func TestUnitShParserErr(t *testing.T) {
	_, err := shload(sheetMeta{})

	if err == nil {
		t.Errorf("an error was expected but was not received")
	}
}

func TestUnitShFindRow(t *testing.T) {
	tn := "TestShFindRow"
	rawData := topen(Mac1621)
	zs, err := zopen(rawData)

	if err != nil {
		t.Errorf("received error from zopen: '%s'", err.Error())
		return
	}

	got := itos(len(zs.info.sheetMeta))
	stmt := "len(zs.info.sheetMeta"
	want := itos(3)
	if got != want {
		t.Errorf(tfail(tn, stmt, got, want))
		return // could panic if not correct
	}

	data, err := shload(zs.info.sheetMeta[0])

	if err != nil {
		t.Error(terr(tn, "data, err := shload(zs.info.sheetMeta[0])", err.Error()))
	}

	next := 0
	openTag := badPair
	closeTag := badPair
	chunk := ""
	isSelfClosing := true

	openTag, isSelfClosing = shFindFirstOccurenceOfElement(data, next, len(data), "row")
	closeTag, isSelfClosing = shTagCloseFind(data, openTag.last+1, len(data), "row")
	stmt = fmt.Sprintf("first, last = shFindFirstOccurenceOfElement(data, %d)", next)
	got = itos(openTag.first)
	want = itos(999)

	if got != want {
		t.Error(tfail(tn, stmt+"; first", got, want))
	}

	got = itos(openTag.last)
	want = itos(1044)

	if got != want {
		t.Error(tfail(tn, stmt+"; last", got, want))
	}

	got = btos(isSelfClosing)
	want = btos(false)

	if got != want {
		t.Error(tfail(tn, "isSelfClosing", got, want))
	}

	if openTag.first < 0 || openTag.first >= len(data) || openTag.last < 0 || openTag.last >= len(data) || openTag.first > openTag.last {
		return // avoid a panic
	}

	chunk = string(data[openTag.first : closeTag.last+1])
	if len(chunk) < 7 {
		t.Error("expected the row we found to have some length to it")
		return // avoid a panic
	}

	got = chunk[0:5]
	stmt = "chunk[0:5]"
	want = "<row "

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}

	got = chunk[len(chunk)-6:]
	stmt = "chunk[len(chunk)-6:]"
	want = "</row>"

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShAdvanceBad(t *testing.T) {
	tn := "TestShAdvanceBad"
	runes := ""
	start := -1
	r := '<'
	expected := -1
	actual := shadvance([]rune(runes), start, r)
	stmt := fmt.Sprintf("shadvance('%s', %d, '%s')", runes, start, string(r))
	got := itos(actual)
	want := itos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShAdvanceGood(t *testing.T) {
	tn := "TestShAdvanceGood"
	runes := " ü Hello Günter"
	start := 2
	r := 'ü'
	expected := 10
	actual := shadvance([]rune(runes), start, r)
	stmt := fmt.Sprintf("shadvance('%s', %d, '%s')", runes, start, string(r))
	got := itos(actual)
	want := itos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShAdvanceNotFound(t *testing.T) {
	tn := "TestShAdvanceGood"
	runes := " ü Hello Günter"
	start := 2
	r := 'x'
	expected := -1
	actual := shadvance([]rune(runes), start, r)
	stmt := fmt.Sprintf("shadvance('%s', %d, '%s')", runes, start, string(r))
	got := itos(actual)
	want := itos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShBadA(t *testing.T) {
	tn := "TestShBadA"
	ix := -1
	runes := "abc"
	expected := true
	actual := shbad([]rune(runes), ix)
	stmt := fmt.Sprintf("shbad('%s', %d)", runes, ix)
	got := btos(actual)
	want := btos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShBadB(t *testing.T) {
	tn := "TestShBadB"
	ix := 3
	runes := "abc"
	expected := true
	actual := shbad([]rune(runes), ix)
	stmt := fmt.Sprintf("shbad('%s', %d)", runes, ix)
	got := btos(actual)
	want := btos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitShBadC(t *testing.T) {
	tn := "TestShBadC"
	ix := 2
	runes := "abc"
	expected := false
	actual := shbad([]rune(runes), ix)
	stmt := fmt.Sprintf("shbad('%s', %d)", runes, ix)
	got := btos(actual)
	want := btos(expected)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

type input struct {
	xml           string
	first         int
	last          int
	tag           string
	open          indexPair
	close         indexPair
	isSelfClosing bool
}

var inputs = []input{
	input{
		xml:           "sg07bloopsgn<jk:bloop >dføg978sg9</     bloop>",
		first:         0,
		last:          -1,
		tag:           "bloop",
		open:          indexPair{12, 22},
		close:         indexPair{33, 45},
		isSelfClosing: false,
	},
	input{
		xml:           "sg07< bloopsgn<jk:bloop >dfsg978sg9<><><><SFG</     bloop>",
		first:         0,
		last:          -1,
		tag:           "bloop",
		open:          badPair,
		close:         badPair,
		isSelfClosing: false,
	},
	input{
		xml:           "<hello:row><row></row></ hello:row>whatever",
		first:         0,
		last:          -1,
		tag:           "row",
		open:          indexPair{0, 10},
		close:         indexPair{22, 34},
		isSelfClosing: false,
	},
	input{
		xml:           "<a><b/></a>",
		first:         0,
		last:          -1,
		tag:           "a",
		open:          indexPair{0, 2},
		close:         indexPair{7, 10},
		isSelfClosing: false,
	},
	input{
		xml:           "<a><b/></a>",
		first:         2,
		last:          -1,
		tag:           "b",
		open:          indexPair{3, 6},
		close:         indexPair{-1, -1}, // shTagCloseFind is not designed to work with self-closing tags
		isSelfClosing: true,
	},
	input{
		xml:           "<row>" + strRowCloseTest,
		first:         0,
		last:          -1,
		tag:           "row",
		open:          indexPair{0, 4},
		close:         indexPair{921, 926}, // shTagCloseFind is not designed to work with self-closing tags
		isSelfClosing: false,
	},
}

func TestUnitShTagFind(t *testing.T) {
	for ix, input := range inputs {
		result, _, isOpenTagSelfClosing := shTagOpenFind([]rune(input.xml), input.first, input.last, input.tag)
		expected := input.open
		if result != expected {
			t.Error(tagFindOpenErr(ix, expected, result, input))
		}

		// note: shTagCloseFind is not designed to work with self-closing tags
		if !isOpenTagSelfClosing {
			start := result.last + 1
			result, _ := shTagCloseFind([]rune(input.xml), start, input.last, input.tag)

			expected = input.close
			if result != expected {
				t.Error(tagFindCloseErr(ix, expected, result, input))
			}
		}

		finalResult, isSelfClosing := shTagFind([]rune(input.xml), input.first, input.last, input.tag)
		finalExpected := tagLoc{open: input.open, close: input.close}

		if input.isSelfClosing {
			selfCloseLoc := indexPair{
				input.open.last,
				input.open.last,
			}
			finalExpected = tagLoc{open: input.open, close: selfCloseLoc}
		}

		if finalResult != finalExpected {
			t.Error(tagFindErr(ix, finalExpected, finalResult, input))
		}

		if isSelfClosing != input.isSelfClosing {
			t.Error(tfail(t.Name(), "isSelfClosing", btos(isSelfClosing), btos(input.isSelfClosing)))
		}
	}
}

func TestUnitShTagNameFind(t *testing.T) {
	tn := "TestShTagNameFind"
	type tagNameFindInput struct {
		str           string
		first         int
		last          int
		expectedElem  string
		expectedLast  int
		isSelfClosing bool
	}

	inputs := []tagNameFindInput{
		{str: "skgshj:cberoy list=\"hi\" >",
			first:         0,
			last:          8888,
			expectedElem:  "cberoy",
			expectedLast:  24,
			isSelfClosing: false,
		},
		{str: "<x/>",
			first:         1,
			last:          8888,
			expectedElem:  "x",
			expectedLast:  3,
			isSelfClosing: true,
		},
		{str: "<x    />",
			first:         1,
			last:          8888,
			expectedElem:  "x",
			expectedLast:  7,
			isSelfClosing: true,
		},
		{str: "<ns:x    />",
			first:         1,
			last:          8888,
			expectedElem:  "x",
			expectedLast:  10,
			isSelfClosing: true,
		},
	}
	for _, input := range inputs {
		elem, last, isSelfClosing := shTagNameFind([]rune(input.str), input.first, input.last)
		stmnt := fmt.Sprintf("elem, last, isSelfClosing := shTagNameFind([]rune(\"%s\"), %d, %d)", input.str, input.first, input.last)

		if elem != input.expectedElem {
			t.Errorf(tfail(tn, stmnt+" -> elem", elem, input.expectedElem))
		}

		if last != input.expectedLast {
			t.Errorf(tfail(tn, stmnt+" -> last", itos(last), itos(input.expectedLast)))
		}

		if isSelfClosing != input.isSelfClosing {
			t.Errorf(tfail(tn, stmnt+" -> isSelfClosing", btos(isSelfClosing), btos(input.isSelfClosing)))
		}
	}
}

func tagFindOpenErr(index int, want, got indexPair, input input) string {
	fn := "shTagOpenFind"
	statement := fmt.Sprintf("input index %d: %s([]rune(\"%s\"), %d, \"%s\")", index, fn, input.xml, input.first, input.tag)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func tagFindCloseErr(index int, want, got indexPair, input input) string {
	fn := "shTagCloseFind"
	statement := fmt.Sprintf("input index %d: %s([]rune(\"%s\"), %d, \"%s\")", index, fn, input.xml, input.first, input.tag)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func tagFindErr(index int, want, got tagLoc, input input) string {
	fn := "shTagFind"
	statement := fmt.Sprintf("input index %d: %s([]rune(\"%s\"), %d, \"%s\")", index, fn, input.xml, input.first, input.tag)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func TestUnitShSetLast(t *testing.T) {
	tn := "TestShSetLast"
	input := -13
	runes := []rune("0123")
	expected := 4
	output := shSetLast(runes, input)
	got := itos(output)
	want := itos(expected)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("shSetLast(\"%s\", %d)", string(runes), input), got, want))
	}

	input = 2
	expected = 2
	output = shSetLast(runes, input)
	got = itos(output)
	want = itos(expected)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("shSetLast(\"%s\", %d)", string(runes), input), got, want))
	}

	input = 3
	expected = 3
	output = shSetLast(runes, input)
	got = itos(output)
	want = itos(expected)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("shSetLast(\"%s\", %d)", string(runes), input), got, want))
	}

	input = 4
	expected = 3
	output = shSetLast(runes, input)
	got = itos(output)
	want = itos(expected)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("shSetLast(\"%s\", %d)", string(runes), input), got, want))
	}
}

// TODO - finish writing this test
//func TestUnitGsheetRowClose(t *testing.T) {
//	str := strRowCloseTest
//	got, isSelfClosing := shTagCloseFind([]rune(str), 0, len(strRowCloseTest), "row")
//	use(got)
//	use(isSelfClosing)
//}
