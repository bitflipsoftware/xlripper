package xmlprivate

import "encoding/xml"

type C struct {
	xml          xml.Name     `xml:"c"`
	R            string       `xml:"r,attr"`
	T            string       `xml:"t,attr"`
	V            string       `xml:"v"`
	InlineString InlineString `xml:"is"`
}

//<c r="A3" s="7" t="inlineStr"><is><t>Everyman</t></is></c>

type InlineString struct {
	xml xml.Name `xml:"is"`
	Str string   `xml:"t"`
}
