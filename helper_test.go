package xlripper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	epsilon          = 0.000001
	testMac1621      = "mac-16.21.xlsx"
	testSharedString = "sharedStrings.xml"
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

func tfail(test, statement, got, want string) string {
	return fmt.Sprintf("test: %s, '%s' = '%s', want '%s'", test, statement, got, want)
}

func terr(test, statement, err string) string {
	return fmt.Sprintf("test: %s, '%s' return an error: '%s'", test, statement, err)
}

func btos(in bool) string {
	return fmt.Sprintf("%t", in)
}

func itos(in int) string {
	return fmt.Sprintf("%d", in)
}

// TODO - delete this
//func TestUnitNothing(t *testing.T) {
//	file, err := os.Open("/Users/mjb/Desktop/qlikview-raw/xl/worksheets/sheet1.xml")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	f, _ := os.Create("/Users/mjb/Desktop/q.xml")
//	f.Write([]byte("hi"))
//	defer f.Close()
//}
