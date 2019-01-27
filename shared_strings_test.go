package xlripper

import "testing"

func TestSharedStringsBlank(t *testing.T) {
	shStrings := newSharedStrings()
	shString := newSharedString()
	str := "hi"
	shString.s = &str
	shStrings.add(shString)

	got := shStrings.get(2)
	want := ""

	if got == nil {
		t.Error("unexpected nil was returned, aborting test")
		return
	}

	if *got != want {
		t.Error(tfail("TestSharedStringsBlank", "shStrings.get(2)", *got, want))
	}
}

func TestSharedStringsNeg(t *testing.T) {
	shStrings := newSharedStrings()
	shString := newSharedString()
	str := "hi"
	shString.s = &str
	shStrings.add(shString)

	got := shStrings.get(-1)
	want := ""

	if got == nil {
		t.Error("unexpected nil was returned, aborting test")
		return
	}

	if *got != want {
		t.Error(tfail("TestSharedStringsNeg", "shStrings.get(-1)", *got, want))
	}
}
