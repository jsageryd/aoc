package main

import (
	"testing"
)

func TestInverseCaptcha(t *testing.T) {
	for n, tc := range []struct {
		in        string
		lookahead int
		out       int
	}{
		// Part 1
		{"1122", 1, 3},
		{"1111", 1, 4},
		{"1234", 1, 0},
		{"91212129", 1, 9},

		// Part 2
		{"1212", 4 / 2, 6},
		{"1221", 4 / 2, 0},
		{"123425", 6 / 2, 4},
		{"123123", 6 / 2, 12},
		{"12131415", 8 / 2, 4},
	} {
		if got, want := inverseCaptcha(tc.in, tc.lookahead), tc.out; got != want {
			t.Errorf("[%d] inverseCaptcha(%q) = %d, want %d", n, tc.in, got, want)
		}
	}
}
