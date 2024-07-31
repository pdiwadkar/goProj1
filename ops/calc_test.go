package ops

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 45
	b := 98
	expected := 143

	if ans := Add(a, b); ans != expected {
		t.Errorf("Add %d %d = %d didn't return %d\n", a, b, ans, expected)
	}
}
func TestMult(t *testing.T) {
	a := 12
	b := 18
	expected := 216

	if ans := Mult(a, b); ans != expected {
		t.Errorf("Mult %d %d = %d didn't return %d\n", a, b, ans, expected)
	}
}
