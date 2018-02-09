package snake

import (
	"testing"
)

func TestNewWhereami(t *testing.T) {
	r := NewWhereami()
	t.Log("Variable content", r)
	if r == nil {
		t.Error("result is not emty")
	}
}
