package xlripper

import (
	"encoding/json"
	"encoding/xml"

	"github.com/bitflip-software/xlripper/xmlprivate"
)

type celLTypeInfo int

const (
	ctUnknown celLTypeInfo = iota
	ctNone
	ctSharedString
	ctInlineString
)

const (
	ctStrUnkown       = "unk"
	ctStrNone         = ""
	ctStrSharedString = "s"
	ctStrInlineString = "inlineStr"
)

func (c celLTypeInfo) String() string {
	switch c {
	case ctUnknown:
		return ctStrUnkown
	case ctNone:
		return ctStrNone
	case ctSharedString:
		return ctStrSharedString
	case ctInlineString:
		return ctStrInlineString
	default:
		return ctStrUnkown
	}

	return ctStrUnkown
}

func (c *celLTypeInfo) Parse(s string) {
	switch s {
	case ctStrUnkown:
		*c = ctUnknown
	case ctStrNone:
		*c = ctNone
	case ctStrSharedString:
		*c = ctSharedString
	case ctStrInlineString:
		*c = ctInlineString
	default:
		*c = ctUnknown
	}
}

type cellCore interface {
	cellReference() string
	cellReferenceRunes() []rune
	typeInfo() celLTypeInfo
	value() *string
	valueRunes() []rune
	json.Marshaler
	parseXML(runes []rune) error
	toXML() ([]rune, error)
}

type cellCoreXML struct {
	x xmlprivate.CellXML
}

func (c *cellCoreXML) cellReference() string {
	return c.x.R
}

func (c *cellCoreXML) cellReferenceRunes() []rune {
	return []rune(c.cellReference())
}

func (c *cellCoreXML) typeInfo() celLTypeInfo {
	if c.x.T == "" {
		return ctNone
	} else if c.x.T == "inlineStr" {
		return ctInlineString
	} else if c.x.T == "s" {
		return ctSharedString
	}

	return ctUnknown
}

func (c *cellCoreXML) value() *string {
	t := c.typeInfo()

	if t == ctInlineString {
		return &c.x.InlineString.Str
	}

	return &c.x.V
}

func (c *cellCoreXML) valueRunes() []rune {
	return []rune(*c.value())
}

func (c *cellCoreXML) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.x)

}

func (c *cellCoreXML) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &c.x)
}

func (c *cellCoreXML) parseXML(runes []rune) error {
	return xml.Unmarshal([]byte(string(runes)), &c.x)
}

func (c *cellCoreXML) toXML() ([]rune, error) {
	b, err := xml.Marshal(c.x)

	if err != nil {
		return nil, err
	}

	return []rune(string(b)), nil
}
