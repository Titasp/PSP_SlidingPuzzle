package input

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler("test input handler")
	if handler == nil {
		t.Fatal("input.NewHandler returned nil")
	}
	t.Log("success")
}
