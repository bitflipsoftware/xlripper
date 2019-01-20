package xlsx

import (
	"testing"
)

func TestZ(t *testing.T) {
	tn := "TestZ"
	raw := topen(Mac1621)
	z, err := zopen(raw)

	if err != nil {
		t.Errorf("en error occurred during zopen: %s", err.Error())
		return
	}

	// check the parsing of the content types file
	got := btos(z.info.contentTypesFound)
	want := btos(true)

	if got != want {
		t.Error(tfail(tn, "z.info.contentTypesFound", got, want))
	}

	got = itos(z.info.contentTypesIndex)
	want = itos(0)

	if got != want {
		t.Error(tfail(tn, "itos(z.info.contentTypesIndex)", got, want))
	}

	got = itos(len(z.info.contentTypes.Defaults))
	want = itos(2)

	if got != want {
		t.Error(tfail(tn, "itos(len(z.info.contentTypes.Defaults))", got, want))
	}

	got = itos(len(z.info.contentTypes.Overrides))
	want = itos(9)

	if got != want {
		t.Error(tfail(tn, "itos(len(z.info.contentTypes.Overrides))", got, want))
	}
}
