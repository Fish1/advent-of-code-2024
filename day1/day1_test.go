package adventofcode2024

import (
	"testing"
)

func TestDay1p1(t *testing.T) {
	want := 11
	have := Day1p1("input/test1.txt")

	if have != want {
		t.Fatalf("have = %d, want = %d", have, want)
	}
}

func TestDay1p2(t *testing.T) {
	want := 31
	have := Day1p2("input/test1.txt")
	if have != want {
		t.Fatalf("have = %d, want = %d", have, want)
	}
}
