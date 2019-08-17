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
