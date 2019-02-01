package xlripper

type cellCoreFast struct {
	r   []rune
	t   celLTypeInfo
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
	if c.ref == badPair {
		return emptyRunes
	}

	return c.r[c.ref.first : c.ref.last+1]
}

func (c *cellCoreFast) typeInfo() celLTypeInfo {
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
	if c.val == badPair {
		return emptyRunes
	}

	return c.r[c.val.first : c.val.last+1]
}

func (c *cellCoreFast) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (c *cellCoreFast) UnmarshalJSON(b []byte) error {
	return nil
}

func (c *cellCoreFast) parseXML(runes []rune) error {
	return nil
}

func (c *cellCoreFast) toXML() ([]rune, error) {
	return nil, nil
}
