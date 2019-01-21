package xlsx

import (
	"fmt"
	"testing"
)

func TestMini(t *testing.T) {
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

func TestMaxi(t *testing.T) {
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
