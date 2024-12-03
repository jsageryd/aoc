package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	if got, want := part1(input), 161; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
