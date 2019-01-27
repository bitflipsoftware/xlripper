package xlripper

type sharedString struct {
	s *string
}

func newSharedString() sharedString {
	return sharedString{
		s: new(string),
	}
}
