package layout

import (
	"strings"
	"testing"
)

func TestNewTile(t *testing.T) {

	testId := "3"

	tile := NewTile(testId)
	if !strings.EqualFold(tile.GetId(), testId) {
		t.Fatal("returned tile id mismatch")
	}
	t.Log("success")
}
