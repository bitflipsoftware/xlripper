package xmlprivate

import "encoding/xml"

type C struct {
	xml xml.Name `xml:"c"`
	R   string   `xml:"r,attr"`
	T   string   `xml:"t,attr"`
	V   string   `xml:"v"`
}
