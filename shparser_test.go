package xlsx

import (
	"fmt"
	"testing"
)

func TestShParserBasics(t *testing.T) {
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

	use(sh)

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
}

func TestShFindRow(t *testing.T) {
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
	tloc := badPair
	chunk := ""

	tloc = shFindFirstOccurenceOfElement(data, next, len(data), "row")
	stmt = fmt.Sprintf("first, last = shFindFirstOccurenceOfElement(data, %d)", next)
	got = itos(tloc.first)
	want = itos(999)

	if got != want {
		t.Error(tfail(tn, stmt+"; first", got, want))
	}

	got = itos(tloc.last)
	want = itos(1044)

	if got != want {
		t.Error(tfail(tn, stmt+"; last", got, want))
	}

	if tloc.first < 0 || tloc.first >= len(data) || tloc.last < 0 || tloc.last >= len(data) || tloc.first > tloc.last {
		return // avoid a panic
	}

	chunk = string(data[tloc.first:tloc.last])
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
	want = "<row>"

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestShAdvanceBad(t *testing.T) {
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

func TestShAdvanceGood(t *testing.T) {
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

func TestShAdvanceNotFound(t *testing.T) {
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

func TestShBadA(t *testing.T) {
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

func TestShBadB(t *testing.T) {
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

func TestShBadC(t *testing.T) {
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
	xml   string
	first int
	last  int
	tag   string
	open  indexPair
	close indexPair
}

var inputs = []input{
	input{
		xml:   "sg07bloopsgn<jk:bloop >dføg978sg9</     bloop>",
		first: 0,
		last:  -1,
		tag:   "bloop",
		open:  indexPair{12, 22},
		close: indexPair{33, 45},
	},
	input{
		xml:   "sg07< bloopsgn<jk:bloop >dfsg978sg9<><><><SFG</     bloop>",
		first: 0,
		last:  -1,
		tag:   "bloop",
		open:  badPair,
		close: badPair,
	},
	input{
		xml:   "<hello:row><row></row></ hello:row>whatever",
		first: 0,
		last:  -1,
		tag:   "row",
		open:  indexPair{0, 10},
		close: indexPair{22, 34},
	},
}

func TestShTagFind(t *testing.T) {

	for ix, input := range inputs {
		result, _ := shTagOpenFind([]rune(input.xml), input.first, input.last, input.tag)
		expected := input.open
		if result != expected {
			t.Error(tagFindOpenErr(ix, expected, result))
		}
		start := result.last + 1
		result = shTagCloseFind([]rune(input.xml), start, input.last, input.tag)
		expected = input.close
		if result != expected {
			t.Error(tagFindCloseErr(ix, expected, result))
		}

		finalResult := shTagFind([]rune(input.xml), input.first, input.last, input.tag)
		finalExpected := tagLoc{open: input.open, close: input.close}
		if finalResult != finalExpected {
			t.Error(tagFindErr(ix, finalExpected, finalResult))
		}
	}
}

func TestShTagNameFind(t *testing.T) {
	tn := "TestShTagNameFind"
	type tagNameFindInput struct {
		str          string
		first        int
		last         int
		expectedElem string
		expectedLast int
	}

	inputs := []tagNameFindInput{
		{str: "skgshj:cberoy list=\"hi\" >",
			first:        0,
			last:         8888,
			expectedElem: "cberoy",
			expectedLast: 24,
		},
	}
	for _, input := range inputs {
		elem, last := shTagNameFind([]rune(input.str), input.first, input.last)

		if elem != input.expectedElem {
			t.Errorf(tfail(tn, "elem, last := shTagNameFind([]rune(input.str), input.first, input.last) -> elem", elem, input.expectedElem))
		}

		if last != input.expectedLast {
			t.Errorf(tfail(tn, "elem, last := shTagNameFind([]rune(input.str), input.first, input.last) -> last", itos(last), itos(input.expectedLast)))
		}
	}
}

func tagFindOpenErr(index int, want, got indexPair) string {
	statement := fmt.Sprintf("input index %d: shTagOpenFind([]rune(input.xml), input.first, input.tag)", index)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func tagFindCloseErr(index int, want, got indexPair) string {
	statement := fmt.Sprintf("input index %d: shTagCloseFind([]rune(input.xml), input.first, input.tag)", index)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func tagFindErr(index int, want, got tagLoc) string {
	statement := fmt.Sprintf("input index %d: shTagFind([]rune(input.xml), input.first, input.tag)", index)
	gots := fmt.Sprintf("%v", got)
	wants := fmt.Sprintf("%v", want)
	return tfail("TestShTagFind", statement, gots, wants)
}

func TestShSetLast(t *testing.T) {
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
