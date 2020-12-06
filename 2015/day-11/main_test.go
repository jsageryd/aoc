package main

import "testing"

func TestNextPassword(t *testing.T) {
	for n, tc := range []struct {
		current string
		next    string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	} {
		if got, want := nextPassword(tc.current), tc.next; got != want {
			t.Errorf("[%d] nextPassword(%q) = %q, want %q", n, tc.current, got, want)
		}
	}
}

func TestValid(t *testing.T) {
	for n, tc := range []struct {
		password string
		valid    bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
	} {
		if got, want := valid(tc.password), tc.valid; got != want {
			t.Errorf("[%d] valid(%q) = %t, want %t", n, tc.password, got, want)
		}
	}
}
