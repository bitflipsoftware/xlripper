package xmlprivate

import "encoding/xml"

type Rels struct {
	XMLName xml.Name `xml:"Relationships"`
	Rels    []Rel    `xml:"Relationship"`
}

type Rel struct {
	ID     string `xml:"Id,attr"`
	Target string `xml:"Target,attr"`
	Type   string `xml:"Type,attr"`
}
