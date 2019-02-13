package xlripper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime/pprof"
	"strings"
	"testing"
	"time"
)

func BenchmarkSmall(b *testing.B) {
	doBench(b, "gsheet.xlsx")
}

func BenchmarkQv(b *testing.B) {
	doBench(b, "qv.xlsx")
}

func BenchmarkLarge(b *testing.B) {
	doBench(b, "large.xlsx")
}

func doBench(b *testing.B, filename string) {
	//fmt.Print("hi")
	dir := thisDir()
	integ := path.Join(dir, "integ", filename)
	p, err := NewParser(integ)

	if err != nil {
		b.Errorf("error ocurred creating parser: %s", err.Error())
		return
	}

	fp, err := os.Create(filename + ".cpu.out")
	checkit(err)
	checkit(pprof.StartCPUProfile(fp))
	start := time.Now()
	//fmt.Print(p.SheetNames())
	for i := 0; i < b.N; i++ {
		sheet, err := p.ParseOne(0)

		if err != nil {
			b.Errorf("error ocurred during parsing: %s", err.Error())
			return
		}

		if len(sheet.Columns) == 0 {
			b.Errorf("empty sheet")
			return
		}
	}

	pprof.StopCPUProfile()
	fp.Close()
	dur := time.Since(start)
	fmt.Print(dur.String() + "\n\n")
	//for _, col := range sheet.Columns {
	//	for _, cell := range col.Cells {
	//		if cell == nil {
	//			fmt.Print("nil")
	//		} else {
	//			fmt.Print(*cell)
	//		}
	//	}
	//}
}

func BenchmarkParseXML(b *testing.B) {
	//s := ReadFile("/Users/mjb/Desktop/sheet1.xml")
	//s = strings.Replace(s, "</x:c><x:c", "</x:c>\n<x:c", -1)
	//s = strings.Replace(s, "</x:row><x:row", "</x:row>\n<x:row", -1)
	//arr := strings.Split(s, "\n")
	//result := make([]string, 0, len(arr))
	//for _, c := range arr {
	//	if len(c) > 20 {
	//		if c[:5] == "<x:c " {
	//			result = append(result, c)
	//		}
	//	}
	//}
	//
	//s = strings.Join(result[0:500000], "\n")
	//checkit(ioutil.WriteFile("/Users/mjb/Desktop/sheet1-with-newlines.xml", []byte(s), constant.FilePermissions))

	contents := iopen("ctest.xml")
	strArr := strings.Split(contents, "\n")
	cells := make([][]rune, 0, len(strArr))

	for _, s := range strArr {
		cells = append(cells, []rune(s))
	}

	fp, err := os.Create("cells.cpu.out")
	checkit(err)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	start := time.Now()
	checkit(pprof.StartCPUProfile(fp))
	for i := 0; i < 100; i++ {
		for _, runes := range cells {
			cf := cellCoreFast{}
			checkit(cf.parseXML(runes))
		}
	}
	pprof.StopCPUProfile()
	fp.Close()
	dur := time.Since(start)
	fmt.Print(dur.String() + "\n\n")

}

func checkit(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func ReadFile(filename string) string {
	f, err := os.Open(filename)
	checkit(err)
	defer f.Close()
	b := bytes.Buffer{}
	w := bufio.NewWriter(&b)
	io.Copy(w, f)
	return string(b.Bytes())
}
