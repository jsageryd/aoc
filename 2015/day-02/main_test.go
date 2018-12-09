package main

import "testing"

func TestPaperNeeded(t *testing.T) {
	for n, tc := range []struct {
		input string
		need  int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	} {
		if got, want := paperNeeded(tc.input), tc.need; got != want {
			t.Errorf("[%d] paperNeeded(%q) = %d, want %d", n, tc.input, got, want)
		}
	}
}

func TestRibbonNeeded(t *testing.T) {
	for n, tc := range []struct {
		input string
		need  int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	} {
		if got, want := ribbonNeeded(tc.input), tc.need; got != want {
			t.Errorf("[%d] ribbonNeeded(%q) = %d, want %d", n, tc.input, got, want)
		}
	}
}
