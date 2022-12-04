package main

import (
	"fmt"
	"testing"
)

func TestFullyOverlappingPairs(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	gotPairs := fullyOverlappingPairs(input)

	wantPairs := []string{
		"2-8,3-7",
		"6-6,4-6",
	}

	if fmt.Sprint(gotPairs) != fmt.Sprint(wantPairs) {
		t.Errorf("got %v, want %v", gotPairs, wantPairs)
	}
}

func TestParsePair(t *testing.T) {
	gotA, gotB := parsePair("1-11,2-22")

	wantA, wantB := rang{1, 11}, rang{2, 22}

	if fmt.Sprint(gotA, gotB) != fmt.Sprint(wantA, wantB) {
		t.Errorf("got %v, %v, want %v, %v", gotA, gotB, wantA, wantB)
	}
}

func TestRang_Contains(t *testing.T) {
	for n, tc := range []struct {
		rangeA, rangeB rang
		contains       bool
	}{
		{rang{2, 8}, rang{3, 7}, true},
		{rang{2, 8}, rang{2, 2}, true},
		{rang{2, 8}, rang{2, 3}, true},
		{rang{2, 8}, rang{8, 8}, true},
		{rang{2, 8}, rang{7, 8}, true},
		{rang{2, 8}, rang{1, 3}, false},
		{rang{2, 8}, rang{7, 9}, false},
		{rang{2, 8}, rang{0, 1}, false},
		{rang{2, 8}, rang{9, 10}, false},
		{rang{2, 8}, rang{1, 1}, false},
		{rang{2, 8}, rang{9, 9}, false},
	} {
		if got, want := tc.rangeA.contains(tc.rangeB), tc.contains; got != want {
			t.Errorf("[%d] %v.contains(%v) = %t, want %t", n, tc.rangeA, tc.rangeB, got, want)
		}
	}
}
