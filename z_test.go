package xlsx

import (
	"fmt"
	"testing"
)

func TestZ(t *testing.T) {
	raw := topen(Mac1621)
	z, err := zopen(raw)

	if err != nil {
		t.Errorf("en error occurred during zopen: %s", err.Error())
		return
	}

	//fmt.Print(z)

	for _, f := range z.r.File {
		if f.FileHeader.Name == "xl/workbook.xml" {
			fmt.Print("hi")
		}
	}
}
