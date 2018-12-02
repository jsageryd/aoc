package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	} {
		if got, want := valid(tc.in), tc.out; got != want {
			t.Errorf("[%d] valid(%q) = %t, want %t", n, tc.in, got, want)
		}
	}
}

func TestValidWithAnagramCheck(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	} {
		if got, want := validWithAnagramCheck(tc.in), tc.out; got != want {
			t.Errorf("[%d] validWithAnagramCheck(%q) = %t, want %t", n, tc.in, got, want)
		}
	}
}

func TestLetterOrderIndependentHash(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out string
	}{
		{"", ""},
		{"abc", "a1b1c1"},
		{"aabbcc", "a2b2c2"},
		{"abcabc", "a2b2c2"},
		{"aaaaaaaaaabc", "a10b1c1"},
	} {
		if got, want := letterOrderIndependentHash(tc.in), tc.out; got != want {
			t.Errorf("[%d] letterOrderIndependentHash(%q) = %q, want %q", n, tc.in, got, want)
		}
	}
}
