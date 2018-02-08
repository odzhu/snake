package snake

import (
	"testing"
)

func TestNewCheck(t *testing.T) {
	r := NewCheck()
	t.Log("Variable content", r)
	if r == nil {
		t.Error("result is not emty")
	}
}

func TestPrintresult(t *testing.T) {
	r := NewCheck()
	r.Name = "test"
	err := r.Printresult()
	if err != nil {
		t.Error("Error")
	}
}
