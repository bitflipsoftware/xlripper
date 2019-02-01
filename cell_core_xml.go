package xlripper

import (
	"encoding/json"
	"encoding/xml"

	"github.com/bitflip-software/xlripper/xmlprivate"
)

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

func (c cellCoreXML) MarshalJSON() ([]byte, error) {
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
