package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	if got, want := part1(input), 157; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDuplicateItem(t *testing.T) {
	for n, tc := range []struct {
		items     string
		duplicate string
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "p"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L"},
		{"PmmdzqPrVvPwwTWBwg", "P"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v"},
		{"ttgJtRGJQctTZtZT", "t"},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "s"},
	} {
		if got, want := duplicateItem(tc.items), tc.duplicate; got != want {
			t.Errorf("[%d] duplicateItem(%q) = %q, want %q", n, tc.items, got, want)
		}
	}
}

func TestSplit(t *testing.T) {
	gotA, gotB := split("vJrwpWtwJgWrhcsFMMfFFhFp")
	wantA, wantB := "vJrwpWtwJgWr", "hcsFMMfFFhFp"

	if gotA != wantA || gotB != wantB {
		t.Errorf("got %q, %q, want %q, %q", gotA, gotB, wantA, wantB)
	}
}

func TestPriority(t *testing.T) {
	for n, tc := range []struct {
		item     string
		priority int
	}{
		{"a", 1},
		{"b", 2},
		{"z", 26},
		{"A", 27},
		{"B", 28},
		{"Z", 52},
	} {
		if got, want := priority(tc.item), tc.priority; got != want {
			t.Errorf("[%d] priority(%q) = %d, want %d", n, tc.item, got, want)
		}
	}
}
