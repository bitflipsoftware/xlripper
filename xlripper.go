package xlripper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Parser struct {
	z zstruct
}

func NewParser(filename string) (Parser, error) {
	ofile, err := os.Open(filename)

	if err != nil {
		return Parser{}, err
	}

	defer ofile.Close()
	buf := bytes.Buffer{}
	bufw := bufio.NewWriter(&buf)
	io.Copy(bufw, ofile)
	return NewParserFromBytes(buf.Bytes())
}

func NewParserFromBytes(b []byte) (Parser, error) {
	p := Parser{}
	z, err := zopen(string(b))

	if err != nil {
		return p, err
	}

	p.z = z
	return p, nil
}

func (p Parser) NumSheets() int {
	return len(p.z.info.sheetMeta)
}

func (p Parser) SheetNames() []string {
	names := make([]string, 0, p.NumSheets())

	for _, x := range p.z.info.sheetMeta {
		names = append(names, x.sheetName)
	}

	return names
}

func (p Parser) ParseOne(sheetIndex int) (Sheet, error) {
	if sheetIndex < 0 {
		return Sheet{}, fmt.Errorf("bad sheet index (negative) %d", sheetIndex)

	}

	if sheetIndex >= p.NumSheets() {
		return Sheet{}, fmt.Errorf("bad sheet index (too large) %d", sheetIndex)
	}

	return shparse(p.z, sheetIndex)
}

func (p Parser) Parse() ([]Sheet, error) {
	sheets := make([]Sheet, 0)
	n := p.NumSheets()

	for i := 0; i < n; i++ {
		sh, err := shparse(p.z, i)

		if err != nil {
			return sheets, err
		}

		sheets = append(sheets, sh)
	}

	return sheets, nil
}
