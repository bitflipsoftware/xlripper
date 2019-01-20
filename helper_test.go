package xlsx

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	Mac1621 = "mac-16.21.xlsx"
)

func topen(filename string) string {
	p := xpath(filename)
	f, err := os.Open(p)

	if err != nil {
		panic(err)
	} else if f == nil {
		panic("file is nil")
	} else {
		defer f.Close()
	}

	buf := bytes.Buffer{}
	w := bufio.NewWriter(&buf)
	io.Copy(w, f)
	return string(buf.Bytes())
}

func xpath(filename string) string {
	dir := testFilesDir()
	rel := filepath.Join(dir, filename)
	abs, err := filepath.Abs(rel)
	if err != nil {
		panic(err)
	}
	return abs
}

func thisFilepath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}

func thisDir() string {
	fp := thisFilepath()
	str := filepath.Dir(fp)
	return str
}

func testFilesDir() string {
	myDir := thisDir()
	rel := filepath.Join(myDir, "test-files")
	abs, err := filepath.Abs(rel)
	if err != nil {
		panic(err)
	}
	return abs
}
