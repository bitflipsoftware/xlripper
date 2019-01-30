package xmlprivate

import (
	"encoding/xml"
)

type C struct {
	xml          xml.Name     `xml:"c" json:"xml_info"`
	R            string       `xml:"r,attr" json:"ref"`
	T            string       `xml:"t,attr" json:"type"`
	V            string       `xml:"v" json:"value"`
	InlineString InlineString `xml:"is" json:"inline_string"`
}

//<c r="A3" s="7" t="inlineStr"><is><t>Everyman</t></is></c>

type InlineString struct {
	xml xml.Name `xml:"is" json:"xml_info"`
	Str string   `xml:"t" json:"string_value"`
}

func ParseXMLC(s string) (C, error) {
	xmlC := C{}
	err := xml.Unmarshal([]byte(s), &xmlC)
	return xmlC, err
}
