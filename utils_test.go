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

func TestRemoveLeadingSlashZeroLen(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := ""
	want := ""
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestRemoveLeadingSlashOnlySlash(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "/"
	want := ""
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestRemoveLeadingSlashOneChar(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "x"
	want := "x"
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

func TestRemoveLeadingSlashRemove(t *testing.T) {
	tn := "TestRemoveLeadingSlashZeroLen"
	input := "/x"
	want := "x"
	got := removeLeadingSlash(input)
	stmt := fmt.Sprintf("removeLeadingSlash(\"%s\")", input)

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}

//func removeLeadingSlash(instr string) (outstr string) {
//	if len(instr) == 0 {
//		return instr
//	} else if len(instr) == 1 && instr == "/" {
//		return ""
//	} else if len(instr) == 1 && instr != "/" {
//		return instr
//	}
//
//	var first rune
//	for _, r := range instr {
//		first = r
//		break
//	}
//
//	if first == '/' {
//		return instr[1:]
//	}
//
//	return instr
//}
