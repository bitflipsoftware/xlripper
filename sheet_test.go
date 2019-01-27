package xlsx

import "testing"

func TestSheet(t *testing.T) {
	sh := NewSheet()
	str := "="
	sh.add(-1, 1, &str)
	sh.add(1, -2, &str)
	got := len(sh.Columns)
	want := 0

	if got != want {
		t.Error(tfail(t.Name(), "len(sh.Columns)", itos(got), itos(want)))
	}
}
