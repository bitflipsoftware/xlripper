package xmlprivate

import "encoding/xml"

type ContentTypeItem struct {
	ContentType string `xml:"ContentType,attr""`
	Extension   string `xml:"Extension,attr"`
	PartName    string `xml:"PartName,attr"`
}

type ContentTypes struct {
	XMLName   xml.Name          `xml:"Types"`
	Defaults  []ContentTypeItem `xml:"Default"`
	Overrides []ContentTypeItem `xml:"Override"`
}
