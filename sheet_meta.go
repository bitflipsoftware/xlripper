package xlsx

import (
	"archive/zip"
	"sort"
)

type sheetMeta struct {
	sheetName  string
	sheetIndex int
	file       *zip.File
	fileIndex  int
	sheetID    string
	relsID     string
}

type sheetMetas []sheetMeta

func (s sheetMetas) sort() {
	sort.Sort(sheetMetas(s))
}

func (s sheetMetas) Len() int {
	return len(s)
}

func (s sheetMetas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sheetMetas) Less(i, j int) bool {
	a := s[i].sheetIndex
	b := s[j].sheetIndex
	return a < b
}
