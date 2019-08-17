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

func TestHasN(t *testing.T) {
	for n, tc := range []struct {
		s    string
		n    int
		set  string
		want bool
	}{
		{"", 0, "", false},
		{"", 1, "", false},
		{"", 2, "", false},
		{"", 3, "", false},
		{"", 3, "aeiou", false},
		{"a", 1, "a", true},
		{"a", 2, "a", false},
		{"aa", 1, "a", true},
		{"aa", 2, "a", true},
		{"aa", 3, "a", false},
		{"a", 3, "aeiou", false},
		{"aa", 3, "aeiou", false},
		{"aaa", 3, "aeiou", true},
		{"eee", 3, "aeiou", true},
		{"iii", 3, "aeiou", true},
		{"ooo", 3, "aeiou", true},
		{"uuu", 3, "aeiou", true},
		{"xxx", 3, "aeiou", false},
		{"aei", 3, "aeiou", true},
		{"eio", 3, "aeiou", true},
		{"iou", 3, "aeiou", true},
		{"a_e_i", 3, "aeiou", true},
		{"_a_e_i", 3, "aeiou", true},
		{"a_e_i_", 3, "aeiou", true},
		{"_a_e_i_", 3, "aeiou", true},
	} {
		if got, want := hasN(tc.s, tc.n, tc.set), tc.want; got != want {
			t.Errorf("[%d] hasN(%q, %d, %s) = %t, want %t", n, tc.s, tc.n, tc.set, got, want)
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
