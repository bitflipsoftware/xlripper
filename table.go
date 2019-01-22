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
