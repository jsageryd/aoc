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

func TestNice2(t *testing.T) {
	for n, tc := range []struct {
		s    string
		nice bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	} {
		if got, want := nice2(tc.s), tc.nice; got != want {
			t.Errorf("[%d] nice2(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}

func TestHasThreeVowels(t *testing.T) {
	for n, tc := range []struct {
		s    string
		want bool
	}{
		{"", false},
		{"a", false},
		{"aa", false},
		{"aaa", true},
		{"eee", true},
		{"iii", true},
		{"ooo", true},
		{"uuu", true},
		{"xxx", false},
		{"aei", true},
		{"eio", true},
		{"iou", true},
		{"a_e_i", true},
		{"_a_e_i", true},
		{"a_e_i_", true},
		{"_a_e_i_", true},
	} {
		if got, want := hasThreeVowels(tc.s), tc.want; got != want {
			t.Errorf("[%d] hasThreeVowels(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}

func TestHasDoubleLetter(t *testing.T) {
	for n, tc := range []struct {
		s    string
		want bool
	}{
		{"", false},
		{"a", false},
		{"aa", true},
		{"aba", false},
		{"bb", true},
		{"_aa", true},
		{"aa_", true},
		{"_aa_", true},
	} {
		if got, want := hasDoubleLetter(tc.s), tc.want; got != want {
			t.Errorf("[%d] hasDoubleLetter(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}

func TestHasForbiddenString(t *testing.T) {
	for n, tc := range []struct {
		s         string
		forbidden []string
		want      bool
	}{
		{"", []string{}, false},
		{"", []string{""}, true},
		{"a", []string{""}, true},
		{"a", []string{"a"}, true},
		{"abc", []string{"ab"}, true},
		{"abc", []string{"bc"}, true},
		{"abc", []string{"ab", "bc"}, true},
		{"abc", []string{"ab", "x"}, true},
		{"abc", []string{"x", "ab"}, true},
		{"abc", []string{"x", "y"}, false},
	} {
		if got, want := hasForbiddenString(tc.s, tc.forbidden), tc.want; got != want {
			t.Errorf("[%d] hasForbiddenString(%q, %q) = %t, want %t", n, tc.s, tc.forbidden, got, want)
		}
	}
}

func TestHasTwoNonOverlappingPairs(t *testing.T) {
	for n, tc := range []struct {
		s    string
		want bool
	}{
		{"", false},
		{"a", false},
		{"aa", false},
		{"aaa", false},
		{"aaaa", true},
		{"abab", true},
		{"abba", false},
		{"aa_aa", true},
		{"ab_ab", true},
		{"_aa_aa", true},
		{"_ab_ab", true},
		{"_abba", false},
	} {
		if got, want := hasTwoNonOverlappingPairs(tc.s), tc.want; got != want {
			t.Errorf("[%d] hasNonOverlappingPair(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}

func TestHasRepeatedLetterWithOneInbetween(t *testing.T) {
	for n, tc := range []struct {
		s    string
		want bool
	}{
		{"", false},
		{"a", false},
		{"aa", false},
		{"aaa", true},
		{"aba", true},
		{"abba", false},
		{"_aa", false},
		{"_aaa", true},
		{"_aba", true},
		{"aa_", false},
		{"aaa_", true},
		{"aba_", true},
		{"_aa_", false},
		{"_aaa_", true},
		{"_aba_", true},
	} {
		if got, want := hasRepeatedLetterWithOneInbetween(tc.s), tc.want; got != want {
			t.Errorf("[%d] hasRepeatedLetterWithOneInbetween(%q) = %t, want %t", n, tc.s, got, want)
		}
	}
}
