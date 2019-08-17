package main

import "testing"

func TestNice(t *testing.T) {
	for n, tc := range []struct {
		s    string
		nice bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	} {
		if got, want := nice(tc.s), tc.nice; got != want {
			t.Errorf("[%d] nice(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}
