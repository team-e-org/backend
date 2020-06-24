package ptr

import (
	"testing"
)

func TestStringPtr(t *testing.T) {
	a := "a"
	b := "a"
	a_ptr := NewString(a)
	b_ptr := NewString(b)
	if a != b {
		t.Fatalf("Strings, a and b should be same")
	}
	if a_ptr == b_ptr {
		t.Fatalf("Pointer of strings, a and b should be different")
	}
	if *a_ptr != *b_ptr {
		t.Fatalf("Strings, a and b should be same")
	}
}

func TestIntPtr(t *testing.T) {
	a := 1234
	b := 1234
	a_ptr := NewInt(a)
	b_ptr := NewInt(b)
	if a != b {
		t.Fatalf("Integers, a and b should be same")
	}
	if a_ptr == b_ptr {
		t.Fatalf("Pointer of integers, a and b should be different")
	}
	if *a_ptr != *b_ptr {
		t.Fatalf("Integers, a and b should be same")
	}
}
