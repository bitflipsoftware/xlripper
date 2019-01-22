package xmlprivate

import "encoding/xml"

type Workbook struct {
	XMLName xml.Name `xml:"workbook"`
	Sheets  Sheets   `xml:"sheets"`
}

type Sheets struct {
	XMLName xml.Name `xml:"sheets"`
	Sheets  []Sheet  `xml:"sheet"`
}

type Sheet struct {
	XMLName xml.Name `xml:"sheet"`
	Name    string   `xml:"name,attr"`
	SheetID string   `xml:"sheetId,attr"`
	RelsID  string   `xml:"id,attr"`
}

func (w *Workbook) FindSheetByRelID(ID string) (sheetIndex int, sheet Sheet) {
	for ix, s := range w.Sheets.Sheets {
		if s.RelsID == ID {
			return ix, s
		}
	}

	return -1, Sheet{}
}
