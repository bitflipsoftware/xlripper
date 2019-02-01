package xlripper

import (
	"fmt"
	"testing"
)

func TestUnitMini(t *testing.T) {
	tn := "TestMini"
	a := -1
	b := 2
	e := a

	got := itos(mini(a, b))
	want := itos(e)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("mini(%d, %d)", a, b), got, want))
	}

	a = 100
	b = 99
	e = b

	got = itos(mini(a, b))
	want = itos(e)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("mini(%d, %d)", a, b), got, want))
	}
}

func TestUnitMaxi(t *testing.T) {
	tn := "TestMaxi"
	a := -1
	b := 2
	e := b

	got := itos(maxi(a, b))
	want := itos(e)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("maxi(%d, %d)", a, b), got, want))
	}

	a = 100
	b = 99
	e = a

	got = itos(maxi(a, b))
	want = itos(e)

	if got != want {
		t.Error(tfail(tn, fmt.Sprintf("maxi(%d, %d)", a, b), got, want))
	}
}

func TestUnitRemoveLeadingSlashZeroLen(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := ""
	want := ""
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitRemoveLeadingSlashOnlySlash(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "/"
	want := ""
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitRemoveLeadingSlashOneChar(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "x"
	want := "x"
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitRemoveLeadingSlashRemove(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "/x"
	want := "x"
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestUnitUseFunction(t *testing.T) {
	x := 1
	use(x)
}
