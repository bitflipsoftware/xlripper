package xlripper

import (
	"fmt"
	"testing"
)

func TestParseRowIndexCellIndex(t *testing.T) {
	tn := "TestParseRowIndexCellIndex"
	type inp struct {
		str   string
		colIX int
		rowIX int
	}

	inputs := []inp{
		{"A1", 0, 0},
		{"B2", 1, 1},
		{"Z3", 25, 2},
		{"AA4", 26, 3},
		{"AB5", 27, 4},
		{"AZ6", 51, 5},
		{"BA7", 52, 6},
		{"BB8", 53, 7},
		{"XFD9", 16383, 8},
		{"WLL10", 15871, 9},
		{"VPF11", 15293, 10},
		{"UEP12", 14341, 11},
		{"ZX13", 699, 12},
		{"ZZ14", 701, 13},
		{"AAA15", 702, 14},
		{"A?16", -1, -1},
		{"Ü17", -1, -1},
		{"UEP123456", 14341, 123455},
	}

	for _, input := range inputs {
		doRowIndexCellIndexTest(t, input.str, input.rowIX, input.colIX, tn)
	}
}

func doRowIndexCellIndexTest(t *testing.T, input string, expectedRowIX, expectedColIX int, testName string) {
	actualRowIX, actualColIX := parseRowIndexCellIndex(input)

	if expectedRowIX != actualRowIX {
		statement := fmt.Sprintf("parseRowIndexCellIndex(\"%s\"), interationIX", input)
		t.Error(tfail(testName, statement, itos(actualRowIX), itos(expectedRowIX)))
	}

	if expectedColIX != actualColIX {
		statement := fmt.Sprintf("parseRowIndexCellIndex(\"%s\"), colIX", input)
		t.Error(tfail(testName, statement, itos(actualColIX), itos(expectedColIX)))
	}
}

func TestParseRowLettersToNum(t *testing.T) {
	tn := "TestParseRowLettersToNum"
	type inp struct {
		str      string
		expected int
	}

	inputs := []inp{
		{"A", 0},
		{"B", 1},
		{"Z", 25},
		{"AA", 26},
		{"AB", 27},
		{"AZ", 51},
		{"BA", 52},
		{"BB", 53},
		{"XFD", 16383},
		{"WLL", 15871},
		{"VPF", 15293},
		{"UEP", 14341},
		{"ZX", 699},
		{"ZZ", 701},
		{"AAA", 702},
		{"A?", -1},
		{"Ü", -1},
	}

	for _, input := range inputs {
		gotI := lettersToNum(input.str)
		wantI := input.expected
		got := itos(gotI)
		want := itos(wantI)
		statement := fmt.Sprintf("lettersToNum(\"%s\")", input.str)

		if got != want {
			t.Error(tfail(tn, statement, got, want))
		}
	}
}
