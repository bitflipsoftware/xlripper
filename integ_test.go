package xlripper

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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
		sheetName := parser.SheetNames()[sheetIX]
		localTestName := fmt.Sprintf("%s %s", tn, sheetName)
		testSheetDataParsing(t, localTestName, sheetIX, parser, test)
	}
}

func testSheetName(t *testing.T, testName string, sheetIndex int, parser Parser, test IntegTest) {
	want := test.Manifest.Sheets[sheetIndex]
	got := parser.SheetNames()[sheetIndex]
	if want != got {
		t.Errorf(testName, fmt.Sprintf("parser.SheetNames()[%d]", sheetIndex), got, want)
	}
}

func testSheetDataParsing(t *testing.T, testName string, sheetIndex int, parser Parser, test IntegTest) {

}
