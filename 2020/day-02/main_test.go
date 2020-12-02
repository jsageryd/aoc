package main

import "testing"

func TestValid(t *testing.T) {
	for n, tc := range []struct {
		min, max int
		letter   byte
		password string
		valid    bool
	}{
		0: {1, 3, 'a', "abcde", true},
		1: {1, 3, 'b', "cdefg", false},
		2: {2, 9, 'c', "ccccccccc", true},
	} {
		if got, want := valid(tc.min, tc.max, tc.letter, tc.password), tc.valid; got != want {
			t.Errorf("[%d] valid(%d, %d, %q, %q) = %t, want %t", n, tc.min, tc.max, tc.letter, tc.password, got, want)
		}
	}
}

func TestValidPart2(t *testing.T) {
	for n, tc := range []struct {
		pos1, pos2 int
		letter     byte
		password   string
		valid      bool
	}{
		0: {1, 3, 'a', "abcde", true},
		1: {1, 3, 'b', "cdefg", false},
		2: {2, 9, 'c', "ccccccccc", false},
	} {
		if got, want := validPart2(tc.pos1, tc.pos2, tc.letter, tc.password), tc.valid; got != want {
			t.Errorf("[%d] validPart2(%d, %d, %q, %q) = %t, want %t", n, tc.pos1, tc.pos2, tc.letter, tc.password, got, want)
		}
	}
}
