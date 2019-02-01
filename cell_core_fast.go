package xlripper

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"unicode"

	"github.com/bitflip-software/xlripper/xmlprivate"
)

type cellCoreFast struct {
	r   []rune
	t   cellTypeInfo
	ref indexPair
	val indexPair
}

func (c *cellCoreFast) cellReference() string {
	if c.ref == badPair {
		return emptyString
	}

	return string(c.cellReferenceRunes())
}

func (c *cellCoreFast) cellReferenceRunes() []rune {
	if c.ref == badPair || c.r == nil || c.ref.first < 0 || c.ref.last >= len(c.r) {
		return emptyRunes
	}

	return c.r[c.ref.first : c.ref.last+1]
}

func (c *cellCoreFast) typeInfo() cellTypeInfo {
	return c.t
}

func (c *cellCoreFast) value() *string {
	if c.val == badPair {
		return &emptyString
	}

	str := string(c.valueRunes())
	return &str
}

func (c *cellCoreFast) valueRunes() []rune {
	if c.val == badPair || c.r == nil || c.val.first < 0 || c.val.last >= len(c.r) {
		return emptyRunes
	}

	return c.r[c.val.first : c.val.last+1]
}

func (c cellCoreFast) MarshalJSON() ([]byte, error) {
	tempx := cellCoreFastToXMLPrivate(c)
	return json.Marshal(tempx)
}

func (c *cellCoreFast) UnmarshalJSON(b []byte) error {
	tempx := xmlprivate.CellXML{}
	err := json.Unmarshal(b, &tempx)

	if err != nil {
		return err
	}

	*c = cellCoreFastFromXMLPrivate(tempx)
	return nil
}

func (c *cellCoreFast) parseXML(runes []rune) error {
	debug := shdebug(runes, 0, 1000)
	use(debug)
	c.r = runes

	ix := 0
	e := len(runes) - 1

	// advance to point to the first character inside the 'c' tag
	ix++

	if ix > e {
		return errors.New("end was reached too soon")
	}

	// advance past any whitespace to the first char of the element name or namespace
	for ix <= e && unicode.IsSpace(runes[ix]) {
		ix++
	}

	if ix > e {
		return errors.New("end was reached too soon")
	}

	// advance past the namespace
	namespaceColon := shFindNamespaceColon(runes, ix, e)
	if namespaceColon > 0 {
		ix = namespaceColon + 1
	}

	// now we should be pointing at an element name 'c'
	if ix > e {
		return errors.New("end was reached too soon")
	}

	if runes[ix] != 'c' {
		return errors.New("wrong element type")
	}

	ix++

	if ix > e {
		return errors.New("end was reached too soon")
	}

	if runes[ix] != ' ' {
		return errors.New("wrong element type")
	}

	attributes, err := shFindAttributes(runes, ix, e)

	if err != nil {
		return err
	}

	for _, a := range attributes {
		if a == badAttribute {
			continue
		} else if a.name == badPair {
			continue
		}

		name := string(runes[a.name.first : a.name.last+1])

		if name == "r" {
			if a.value != badPair && a.value.last > a.value.first {
				c.ref = a.value
			} else {
				c.ref = badPair
			}
		} else if name == "t" {
			val := ""

			if a.value != badPair && a.value.last > a.value.first {
				val = string(runes[a.value.first : a.value.last+1])
			}

			c.t.Parse(val)
		} else if name == "s" {
			// TODO - what is 's'? What does it mean/represent?
			if a.value != badPair && a.value.last > a.value.first {
				c.val = a.value
			} else {
				c.val = badPair
			}
		}
	}

	if false {
		b := []byte(string(runes))
		tempx := xmlprivate.CellXML{}
		err := xml.Unmarshal(b, &tempx)

		if err != nil {
			return err
		}

		*c = cellCoreFastFromXMLPrivate(tempx)
	}
	return nil
}

func (c *cellCoreFast) toXML() ([]rune, error) {
	tempx := cellCoreFastToXMLPrivate(*c)
	raw, err := xml.Marshal(tempx)

	if err != nil {
		return nil, err
	}

	return []rune(string(raw)), nil
}

func cellCoreFastToXMLPrivate(c cellCoreFast) xmlprivate.CellXML {
	tempx := xmlprivate.CellXML{}
	tempx.R = c.cellReference()
	tempx.T = c.typeInfo().String()

	if c.typeInfo() == ctInlineString {
		tempx.InlineString.Str = *c.value()
	} else {
		tempx.V = *c.value()
	}

	return tempx
}

func cellCoreFastFromXMLPrivate(x xmlprivate.CellXML) cellCoreFast {
	c := cellCoreFast{}
	c.r = make([]rune, 0)
	if len(x.R) > 0 {
		rrunes := []rune(x.R)
		c.ref.first = len(c.r)
		c.ref.last = c.val.first + len(rrunes) - 1
		c.r = append(c.r, rrunes...)
	} else {
		c.ref = badPair
	}

	c.t.Parse(x.T)
	valstr := ""
	if c.typeInfo() == ctInlineString {
		valstr = x.InlineString.Str
	} else {
		valstr = x.V
	}

	if len(valstr) > 0 {
		vrunes := []rune(valstr)
		c.val.first = len(c.r)
		c.val.last = c.val.first + len(vrunes) - 1
		c.r = append(c.r, vrunes...)
	} else {
		c.ref = badPair
	}

	return c
}
