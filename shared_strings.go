package xlripper

type sharedStrings struct {
	sstrings []sharedString
}

func newSharedStrings() sharedStrings {
	return sharedStrings{
		sstrings: make([]sharedString, 0, 100),
	}
}

func (s *sharedStrings) add(shs sharedString) {
	s.sstrings = append(s.sstrings, shs)
}

func (s *sharedStrings) get(ix int) *string {
	if ix < 0 || ix >= len(s.sstrings) {
		blank := ""
		return &blank
	}

	return s.sstrings[ix].s
}

func (s *sharedStrings) len() int {
	return len(s.sstrings)
}
