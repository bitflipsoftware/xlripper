package xlripper

import (
	"path"
	"testing"
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

	//fmt.Print(p.SheetNames())
	sheet, err := p.ParseOne(0)

	if err != nil {
		b.Errorf("error ocurred during parsing: %s", err.Error())
		return
	}

	if len(sheet.Columns) == 0 {
		b.Errorf("empty sheet")
		return
	}

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
