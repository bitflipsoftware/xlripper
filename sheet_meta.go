package xlsx

import "archive/zip"

type sheetMeta struct {
	name      string
	index     int
	file      *zip.File
	fileIndex int
}
