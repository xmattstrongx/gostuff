package sudoku2

import "testing"

func TimeParseHappyPath(t *testing.T) {
	if have, want := "1", "1"; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}
