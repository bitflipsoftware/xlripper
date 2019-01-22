package xlsx

import (
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
