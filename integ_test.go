package xlripper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	extMeta = "meta.json"
)

type Meta struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsFailureExpected bool   `json:"is_failure_expected"`
	Sheets            string `json:"sheets"`
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

metaParseLoop:
	for _, jsonFile := range jsonFiles {
		jsonPath := path.Join(integDir, jsonFile.Name())
		ofile, err := os.Open(jsonPath)

		if err != nil {
			t.Errorf("error trying to open %s - %s", jsonPath, err.Error())
			continue metaParseLoop
		}

		jsonBytes := make([]byte, 0)
		jsonLen, err := ofile.Read(jsonBytes)

		if err != nil {
			t.Errorf("error trying to read bytes of %s - %s", jsonPath, err.Error())
			ofile.Close()
			continue metaParseLoop
		} else if jsonLen == 0 {
			t.Errorf("%s seems to be empty", jsonPath)
			ofile.Close()
			continue metaParseLoop
		}

		ofile.Close()
		m := Meta{}
		err = json.Unmarshal(jsonBytes, &m)

		if err != nil {
			t.Errorf("error unmarshaling %s - %s", jsonPath, err.Error())
			continue metaParseLoop
		}

		use(m)
	}

	fmt.Print(files)
}
