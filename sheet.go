package xlripper

type Sheet struct {
	Name    string
	Index   int
	Columns []Column
	strs    map[string]*string
}

func NewSheet() Sheet {
	return Sheet{
		Name:    "",
		Index:   -1,
		Columns: make([]Column, 0),
		strs:    make(map[string]*string, 100000),
	}
}

func (s *Sheet) add(rowIX, colIX int, val *string) {

	if colIX < 0 {
		return
	}

	if rowIX < 0 {
		return
	}

	for i := len(s.Columns); i <= colIX; i++ {
		s.Columns = append(s.Columns, NewColumn())
	}

	col := s.Columns[colIX]

	for i := len(col.Cells); i <= rowIX; i++ {
		col.Cells = append(col.Cells, nil)
	}

	if val != nil {
		if _, ok := s.strs[*val]; !ok {
			s.strs[*val] = val
		}

		val = s.strs[*val]
	}

	col.Cells[rowIX] = val
	s.Columns[colIX] = col
}
