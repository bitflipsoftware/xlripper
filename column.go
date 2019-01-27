package xlripper

// Column represents a column of values in an xlsx spreadsheet. Cell values are represented as strings.
// Type and formatting information from the spreadsheet is discarded, only a string representation of
// the value remains.
//
// The strings are held as pointers for the sake of memory optimization. You should not mutate these
// as you may be surprised by the results if other columns or cells are pointing to the same string.
// The data structure is intended to be used as a read-only data structure.
//
type Column struct {
	Cells []*string
}

func NewColumn() Column {
	c := Column{}
	c.Cells = make([]*string, 0)
	return c
}
