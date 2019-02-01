package xlripper

import "encoding/json"

type celLTypeInfo int

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
