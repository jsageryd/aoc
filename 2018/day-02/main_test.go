package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	for n, tc := range []struct {
		in  []string
		out int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	} {
		if got, want := checksum(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestCommonLettersOfCorrectBoxes(t *testing.T) {
	for n, tc := range []struct {
		in  []string
		out string
	}{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	} {
		if got, want := commonLettersOfCorrectBoxes(tc.in), tc.out; got != want {
			t.Errorf("[%d] got %q, want %q", n, got, want)
		}
	}
}
