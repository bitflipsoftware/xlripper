package xlripper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

const (
	epsilon = 0.000001
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
func XNothing(t *testing.T) {
	file, err := os.Open("/Users/mjb/Desktop/qlikview-raw/xl/worksheets/sheet1.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bytes.Buffer{}
	wr := bufio.NewWriter(&buf)
	io.Copy(wr, file)
	str := string(buf.Bytes())
	//str = strings.Replace(str, "><", ">\n<", -1)
	use(str)

	//<t>Originating Customer</t>
	//<t>Destination NPA-NXX</t>
	//<t>Jurisdiction State</t>
	//<t>Day (MM/DD/YY (day))</t>
	//<t>Variable Cost</t>
	//<t>Orig Per Minute Revenue</t>
	//<t>Total Minutes</t>
	//<t>Variable Margin</t>
	//<t>Peerless Network Inc</t>
	//<t>Intra-State</t>
	//<t>Inter-State</t>
	//<t>Unidentified</t>
	str = strings.Replace(str, "Originating Customer", "Purple Dragons", -1)
	str = strings.Replace(str, "Destination NPA-NXX", "Count of Headaches", -1)
	str = strings.Replace(str, "Jurisdiction State", "Type of Food", -1)
	//str = strings.Replace(str, "Day (MM/DD/YY (day))", "", -1)
	str = strings.Replace(str, "Variable Cost", "Weight of Balloons", -1)
	str = strings.Replace(str, "Orig Per Minute Revenue", "Breakfast Cost", -1)
	str = strings.Replace(str, "Total Minutes", "Number of Hotdogs", -1)
	str = strings.Replace(str, "Peerless Network Inc", "Everyman", -1)
	str = strings.Replace(str, "Intra-State", "Eggs", -1)
	str = strings.Replace(str, "Inter-State", "Kale", -1)
	str = strings.Replace(str, "Unidentified", "Hamburger", -1)
	str = strings.Replace(str, "", "", -1)
	str = strings.Replace(str, "", "", -1)
	str = strings.Replace(str, "", "", -1)

	f, _ := os.Create("/Users/mjb/Desktop/q.xml")
	f.Write([]byte(str))
	defer f.Close()
}
