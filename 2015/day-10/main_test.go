package main

import "testing"

func TestLookAndSay(t *testing.T) {
	for n, tc := range []struct {
		input  string
		output string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	} {
		if got, want := lookAndSay(tc.input), tc.output; got != want {
			t.Errorf("[%d] lookAndSay(%q) = %q, want %q", n, tc.input, got, want)
		}
	}
}
