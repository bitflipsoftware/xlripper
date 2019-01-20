package xlsx

import (
	"testing"
)

func TestNothing(t *testing.T) {
	str := xopen(Mac1621)
	unzip(str, "/Users/mjb/Desktop")
}
