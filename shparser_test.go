package xlsx

import (
	"fmt"
	"testing"
)

func TestShParserBasics(t *testing.T) {
	tn := "TestShParserBasics"
	rawData := topen(Mac1621)
	zs, err := zopen(rawData)

	if err != nil {
		t.Errorf("received error from zopen: '%s'", err.Error())
		return
	}

	tbl, err := shparse(zs, 0)

	if err != nil {
		t.Errorf("received error from shparse: '%s'", err.Error())
		return
	}

	use(zs)
	use(tn)
	use(tbl)
}

func TestShFindRow(t *testing.T) {
	tn := "TestShFindRow"
	rawData := topen(Mac1621)
	zs, err := zopen(rawData)

	if err != nil {
		t.Errorf("received error from zopen: '%s'", err.Error())
		return
	}

	got := itos(len(zs.info.sheetMeta))
	stmt := "len(zs.info.sheetMeta"
	want := itos(3)
	if got != want {
		t.Errorf(tfail(tn, stmt, got, want))
		return // could panic if not correct
	}

	data, err := shload(zs.info.sheetMeta[0])

	if err != nil {
		t.Error(terr(tn, "data, err := shload(zs.info.sheetMeta[0])", err.Error()))
	}

	next := 0
	first := -1
	last := -1
	chunk := ""

	first, last = shfindRow(data, next)
	stmt = fmt.Sprintf("first, last = shfindRow(data, %d)", next)
	got = itos(first)
	want = itos(999)

	if got != want {
		t.Error(tfail(tn, stmt+"; first", got, want))
	}

	got = itos(last)
	want = itos(1273)

	if got != want {
		t.Error(tfail(tn, stmt+"; last", got, want))
	}

	if first < 0 || first >= len(data) || last < 0 || last >= len(data) || first > last {
		return // avoid a panic
	}

	chunk = string(data[first:last])
	if len(chunk) < 7 {
		return // avoid a panic
	}

	got = chunk[0:5]
	stmt = "chunk[0:5]"
	want = "<row>"

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}

	got = chunk[len(chunk)-6:]
	stmt = "chunk[len(chunk)-6:]"
	want = "<row>"

	if got != want {
		t.Error(tfail(tn, stmt, got, want))
	}
}
