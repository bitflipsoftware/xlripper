package xlsx

import (
	"testing"
)

func TestZBad(t *testing.T) {
	tn := "TestZBad"
	raw := "this is not a valid zip file"
	_, err := zopen(raw)

	got := btos(err == nil)
	want := btos(false)

	if got != want {
		t.Error(tfail(tn, "err == nil", got, want))
	}
}

func TestZ(t *testing.T) {
	tn := "TestZ"
	raw := topen(Mac1621)
	z, err := zopen(raw)

	if err != nil {
		t.Errorf("en error occurred during zopen: %sstrings", err.Error())
		return
	}

	// check the parsing of the content types file
	got := itos(z.info.contentTypesIndex)
	want := itos(0)

	if got != want {
		t.Error(tfail(tn, "z.info.contentTypesIndex", got, want))
	}

	got = itos(len(z.info.contentTypes.Defaults))
	want = itos(2)

	if got != want {
		t.Error(tfail(tn, "len(z.info.contentTypes.Defaults", got, want))
	}

	got = itos(len(z.info.contentTypes.Overrides))
	want = itos(9)

	if got != want {
		t.Error(tfail(tn, "len(z.info.contentTypes.Overrides", got, want))
	}

	// check the parsing of the rels file
	got = itos(z.info.relsIndex)
	want = itos(1)

	if got != want {
		t.Error(tfail(tn, "z.info.relsIndex", got, want))
	}

	got = itos(len(z.info.rels.Rels))
	want = itos(3)

	if got != want {
		t.Error(tfail(tn, "z.info.rels.Rels", got, want))
		return // we might panic if this is not correct
	}

	got = z.info.rels.Rels[2].Target
	want = "xl/workbook.xml"

	if got != want {
		t.Error(tfail(tn, "z.info.rels.Rels[2].Target", got, want))
	}

	got = z.info.rels.Rels[2].ID
	want = "rId1"

	if got != want {
		t.Error(tfail(tn, "z.info.rels.Rels[2].ID", got, want))
	}

	got = z.info.rels.Rels[2].Type
	want = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"

	if got != want {
		t.Error(tfail(tn, "z.info.rels.Rels[2].Type", got, want))
	}

	// see if the workbook was found
	got = btos(z.info.wkbk == nil)
	want = btos(false)

	if got != want {
		t.Error(tfail(tn, "z.info.workbook == nil", got, want))
	}

	got = z.info.wkbkName
	want = "xl/workbook.xml"

	if got != want {
		t.Error(tfail(tn, "z.info.workbookName", got, want))
	}

	got = itos(z.info.sharedStrings.len())
	want = itos(13)

	if got != want {
		t.Error(tfail(tn, "z.info.sharedStrings.len()", got, want))
	}

	// check the shared strings
	got = itos(z.info.wkbkIndex)
	want = itos(3)

	if got != want {
		t.Error(tfail(tn, "z.info.workbookIndex", got, want))
	}

	shs := []string{
		"sfd",
		"sad;fghjk",
		"asfgaf",
		"asdfadsg",
		"asdgg",
		"asfgasdfasdf",
		"asfgadfgadfg",
		"awer5aw5w45",
		"z",
		"b",
		"d",
		"xx",
		"Maybe this will be a shared string.",
	}

	for ix, str := range shs {
		got = *(z.info.sharedStrings.get(ix))
		want = str

		if got != want {
			t.Error(tfail(tn, "z.info.sharedStrings.get(ix)", got, want))
		}
	}
}
