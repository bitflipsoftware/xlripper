package xlsx

import (
	"fmt"
	"testing"
)

func TestParseRowIndexCellIndex(t *testing.T) {
	tn := "TestParseRowIndexCellIndex"
	input := "A1"
	expectedRowIX := 0
	expectedColIX := 0
	doRowIndexCellIndexTest(t, input, expectedRowIX, expectedColIX, tn)
}

func doRowIndexCellIndexTest(t *testing.T, input string, expectedRowIX, expectedColIX int, testName string) {
	actualRowIX, actualColIX := parseRowIndexCellIndex(input)

	if expectedRowIX != actualRowIX {
		statement := fmt.Sprintf("parseRowIndexCellIndex(\"%s\"), rowIX", input)
		t.Error(tfail(testName, statement, itos(actualRowIX), itos(expectedRowIX)))
	}

	if expectedColIX != actualColIX {
		statement := fmt.Sprintf("parseRowIndexCellIndex(\"%s\"), colIX", input)
		t.Error(tfail(testName, statement, itos(actualColIX), itos(expectedColIX)))
	}
}

func TestParseRowLettersToNum(t *testing.T) {
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
		{"ZAUT", 4359876},
	}

	tn := "TestParseRowLettersToNum"
	input := "A1"
	expectedRowIX := 0
	expectedColIX := 0
	doRowIndexCellIndexTest(t, input, expectedRowIX, expectedColIX, tn)
}
