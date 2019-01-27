package xlsx

type Sheet struct {
	Name    string
	Index   int
	Columns []Column
}

func NewSheet() Sheet {
	return Sheet{
		Name:    "",
		Index:   -1,
		Columns: make([]Column, 0),
	}
}

func (s *Sheet) add(rowIX, cellIX int, val *string) {

	if cellIX < 0 {
		return
	}

	if rowIX < 0 {
		return
	}

	for i := len(s.Columns); i < cellIX; i++ {
		s.Columns = append(s.Columns, NewColumn())
	}

	col := s.Columns[cellIX]

	for i := len(col.Cells); i < rowIX; i++ {
		col.Cells = append(col.Cells, nil)
	}

	col.Cells[rowIX] = val
}
