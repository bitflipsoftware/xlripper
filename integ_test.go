package xlripper

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

const (
	extMeta = "meta.json"
	extXLSX = "xlsx"
)

type Manifest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	IsFailureExpected bool     `json:"is_failure_expected"`
	Sheets            []string `json:"sheets"`
}

type IntegTest struct {
	Manifest     Manifest
	MetaFilename string
	Root         string
	FileXLSX     os.FileInfo
	FileSheets   []os.FileInfo
}

func TestInteg(t *testing.T) {
	myDir := thisDir()
	integDir := path.Join(myDir, "integ")

	stats, err := os.Stat(integDir)

	if os.IsNotExist(err) {
		// abort if the git submodule is not there
		fmt.Printf("integ tests are not available\n")
		return
	} else if err != nil {
		fmt.Printf("error inspecting integ dir %s\n", err.Error())
		return
	}
	mode := stats.Mode()

	if !mode.IsDir() {
		fmt.Printf("integ tests are not available\n")
		return
	}

	files, err := ioutil.ReadDir(integDir)

	if err != nil {
		t.Errorf("error listing integ directory %s", err.Error())
	}

	jsonFiles := make([]os.FileInfo, 0)

	for _, f := range files {
		nm := strings.ToLower(f.Name())
		if len(nm) >= len(extMeta) {
			ext := nm[len(nm)-len(extMeta):]
			if ext == extMeta {
				jsonFiles = append(jsonFiles, f)
			}
		}
	}

	tests := make([]IntegTest, 0)

metaParseLoop:
	for _, jsonFile := range jsonFiles {
		jsonPath := path.Join(integDir, jsonFile.Name())
		ofile, err := os.Open(jsonPath)

		if err != nil {
			t.Errorf("error trying to open %s - %s", jsonPath, err.Error())
			continue metaParseLoop
		}

		jsonBuf := bytes.Buffer{}
		jsonBufw := bufio.NewWriter(&jsonBuf)
		io.Copy(jsonBufw, ofile)
		jsonBytes := jsonBuf.Bytes()
		ofile.Close()
		m := Manifest{}
		err = json.Unmarshal(jsonBytes, &m)

		if err != nil {
			t.Errorf("error unmarshaling %s - %s", jsonPath, err.Error())
			continue metaParseLoop
		}

		if len(m.Sheets) == 0 {
			t.Errorf("no sheets were found in the manifest %s", jsonPath)
			continue metaParseLoop
		}

		itest := IntegTest{}
		itest.Manifest = m
		itest.Root = jsonFile.Name()[:len(jsonFile.Name())-len(extMeta)]
		itest.MetaFilename = jsonFile.Name()
		if len(itest.Root) > 0 {
			if itest.Root[0:1] == "." {
				itest.Root = itest.Root[1:]
			}
		}

		if len(itest.Root) > 0 {
			if itest.Root[len(itest.Root)-1:] == "." {
				itest.Root = itest.Root[:len(itest.Root)-1]
			}
		}

		if len(itest.Root) > 0 {
			tests = append(tests, itest)
		}
	}

	tests = gatherTests(t, tests, files, integDir)

	for _, itest := range tests {
		runIntegTest(t, itest, integDir)
	}
}

func gatherTests(t *testing.T, tests []IntegTest, files []os.FileInfo, dir string) []IntegTest {
	for itestIX, itest := range tests {
		expectedX := strings.ToLower(itest.Root + "." + extXLSX)
		if itest.Manifest.Name != expectedX {
			t.Errorf("the name found in %s should be %s", itest.MetaFilename, expectedX)
		}

		isXLSXFound := false

	xlsxFileSearchLoop:
		for _, f := range files {
			if strings.ToLower(f.Name()) == expectedX {
				itest.FileXLSX = f
				isXLSXFound = true
				break xlsxFileSearchLoop
			}
		}

		if !isXLSXFound {
			t.Errorf("the xlsx file %s could not be found", expectedX)
		}

		expectedNumSheets := len(itest.Manifest.Sheets)

		if expectedNumSheets == 0 {
			t.Errorf("the manifest has zero sheets: %s", itest.MetaFilename)
		}

		itest.FileSheets = make([]os.FileInfo, 0)

		for sheetIX, _ := range itest.Manifest.Sheets {
			sheetCsvName := fmt.Sprintf("%s.sheet%d.csv", itest.Root, sheetIX)
			isFound := false
		sheetFileSearchLoop:
			for _, f := range files {
				searchName := strings.ToLower(f.Name())
				if searchName == sheetCsvName {
					itest.FileSheets = append(itest.FileSheets, f)
					isFound = true
					break sheetFileSearchLoop
				}
			}

			if !isFound {
				t.Errorf("%s was not found", sheetCsvName)
			}
		}

		tests[itestIX] = itest
	}

	return tests
}

func runIntegTest(t *testing.T, test IntegTest, dir string) {
	tn := fmt.Sprintf("Integ Test %s", test.FileXLSX.Name())
	xlsxPath := path.Join(dir, test.FileXLSX.Name())
	parser, err := NewParser(xlsxPath)

	if err != nil {
		t.Errorf("error opening parser for %s: %s", test.FileXLSX.Name(), err.Error())
		return
	}

	wantNumSheets := len(test.Manifest.Sheets)
	gotNumSheets := parser.NumSheets()

	if wantNumSheets != gotNumSheets {
		t.Error(tfail(tn, "parser.NumSheets()", itos(gotNumSheets), itos(wantNumSheets)))
		return
	}

	for sheetIX := 0; sheetIX < wantNumSheets; sheetIX++ {
		testSheetName(t, tn, sheetIX, parser, test)
	}

	for sheetIX := 0; sheetIX < wantNumSheets; sheetIX++ {
		//sheetName := parser.SheetNames()[sheetIX]
		localTestName := fmt.Sprintf("%s sheet[%d]", tn, sheetIX)
		testSheetDataParsing(t, localTestName, sheetIX, parser, test, dir)
	}
}

func testSheetName(t *testing.T, testName string, sheetIndex int, parser Parser, test IntegTest) {
	want := test.Manifest.Sheets[sheetIndex]
	got := parser.SheetNames()[sheetIndex]
	if want != got {
		t.Errorf(testName, fmt.Sprintf("parser.SheetNames()[%d]", sheetIndex), got, want)
	}
}

func testSheetDataParsing(t *testing.T, testName string, sheetIndex int, parser Parser, test IntegTest, dir string) {
	//sheetName := parser.SheetNames()[sheetIndex]
	csvFilename := test.FileSheets[sheetIndex].Name()
	csvPath := path.Join(dir, csvFilename)
	ofile, err := os.Open(csvPath)

	if err != nil {
		t.Errorf("error occurred opening %s: %s", csvFilename, err.Error())
	}

	defer ofile.Close()
	csvReader := csv.NewReader(ofile)

	// Read reads one record (a slice of fields) from r. If the record has an unexpected number of fields, Read returns
	// the record along with the error ErrFieldCount. Except for that case, Read always returns either a non-nil record
	// or a non-nil error, but not both. If there is no data left to be read, Read returns nil, io.EOF. If ReuseRecord
	// is true, the returned slice may be shared between multiple calls to Read.

	csvRows := make([][]string, 0)

csvLoop:
	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			break csvLoop
		} else if err != nil {
			t.Errorf("failure parsing %s: %s", csvFilename, err.Error())
			break csvLoop
		}

		csvRows = append(csvRows, row)
	}

	sheet, err := parser.ParseOne(sheetIndex)

	if err != nil {
		t.Errorf("%s: received error during xlsx paring of sheet %d: %s", testName, sheetIndex, err.Error())
	}

	numRows, numCols := findMaxRowAndColumnLengths(sheet, csvRows)

	for rowIX := 0; rowIX < numRows; rowIX++ {
		for colIX := 0; colIX < numCols; colIX++ {
			csvVal, csvOK := getCsvValue(csvRows, rowIX, colIX)
			xlsxVal, xlsxOK := getXLSXValue(sheet, rowIX, colIX)
			thisTest := fmt.Sprintf("%s rowIX %d, colIX %d", testName, rowIX, colIX)

			if csvOK && len(csvVal) > 0 {
				if !xlsxOK {
					t.Errorf(tfail(thisTest, "csvOK && len(csvVal) > 0 && !xlsxOK", btos(true), btos(false)))
				}
			} else if xlsxOK && len(xlsxVal) > 0 {
				if !csvOK {
					t.Errorf(tfail(thisTest, "xlsxOK && len(xlsxVal) > 0 && !csvOK", btos(true), btos(false)))
				}
			}

			testPasses := areEqual(csvVal, xlsxVal)

			if !testPasses {
				// bizarre leading junk seems to be added at the start of a file exported with microsoft excel
				if rowIX == 0 && colIX == 0 {
					csvRunes := []rune(csvVal)
					xlsxRunes := []rune(xlsxVal)

					if len(csvRunes) == len(xlsxRunes)+1 {
						testPasses = areEqual(xlsxVal, string(csvRunes[1:]))
					}
				}
			}

			if !testPasses {
				t.Error(tfail(thisTest, "xlsxVal", xlsxVal, csvVal))
			}
		}
	}
}

func areEqual(want, got string) (equal bool) {
	if want == got {
		return true
	}

	// check if they are different by some floating point imprecision
	csvFloat, csvFloatErr := strconv.ParseFloat(want, 64)
	xlsxFloat, xlsxFloatErr := strconv.ParseFloat(got, 64)

	if csvFloatErr == nil && xlsxFloatErr == nil {
		diff := math.Abs(csvFloat - xlsxFloat)

		if diff < epsilon {
			return true
		}
	}

	return false
}

func findMaxRowAndColumnLengths(sheet Sheet, csvRows [][]string) (numRows, numCols int) {
	xNumCols := len(sheet.Columns)
	xNumRows := -1

	for _, c := range sheet.Columns {
		if len(c.Cells) > xNumRows {
			xNumRows = len(c.Cells)
		}
	}

	cNumRows := len(csvRows)
	cNumCols := -1

	for _, r := range csvRows {
		if len(r) > cNumCols {
			cNumCols = len(r)
		}
	}

	maxCols := maxi(xNumCols, cNumCols)
	maxRows := maxi(xNumRows, cNumRows)
	return maxRows, maxCols

}

func getCsvValue(csvRows [][]string, rowIX, colIX int) (value string, ok bool) {
	if rowIX < 0 || colIX < 0 {
		return "", false
	}

	if rowIX >= len(csvRows) {
		return "", false
	}

	row := csvRows[rowIX]

	if colIX >= len(row) {
		return "", false
	}

	return row[colIX], true
}

func getXLSXValue(sheet Sheet, rowIX, colIX int) (value string, ok bool) {
	if colIX < 0 || rowIX < 0 {
		return "", false
	}

	if colIX >= len(sheet.Columns) {
		return "", false
	}

	col := sheet.Columns[colIX]

	if rowIX >= len(col.Cells) {
		return "", false
	}

	cell := col.Cells[rowIX]

	if cell == nil {
		return "", true
	}

	return *cell, true
}
