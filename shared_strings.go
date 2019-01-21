package xlsx

type sharedStrings struct {
	s []sharedString
}

func newSharedStrings() sharedStrings {
	return sharedStrings{
		s: make([]sharedString, 0, 100),
	}
}

func (s *sharedStrings) add(shs sharedString) {
	s.s = append(s.s, shs)
}
