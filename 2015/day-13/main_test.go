package main

import (
	"fmt"
	"testing"
)

func TestOptimalTotalHappiness(t *testing.T) {
	input := []string{
		"Alice would gain 54 happiness units by sitting next to Bob.",
		"Alice would lose 79 happiness units by sitting next to Carol.",
		"Alice would lose 2 happiness units by sitting next to David.",
		"Bob would gain 83 happiness units by sitting next to Alice.",
		"Bob would lose 7 happiness units by sitting next to Carol.",
		"Bob would lose 63 happiness units by sitting next to David.",
		"Carol would lose 62 happiness units by sitting next to Alice.",
		"Carol would gain 60 happiness units by sitting next to Bob.",
		"Carol would gain 55 happiness units by sitting next to David.",
		"David would gain 46 happiness units by sitting next to Alice.",
		"David would lose 7 happiness units by sitting next to Bob.",
		"David would gain 41 happiness units by sitting next to Carol.",
	}

	attendees, happiness := parse(input)

	if got, want := optimalTotalHappiness(attendees, happiness), 330; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPerm(t *testing.T) {
	perms := [][]string{
		{"a", "b", "c", "d"},
		{"a", "b", "d", "c"},
		{"a", "c", "b", "d"},
		{"a", "c", "d", "b"},
		{"a", "d", "b", "c"},
		{"a", "d", "c", "b"},
		{"b", "a", "c", "d"},
		{"b", "a", "d", "c"},
		{"b", "c", "a", "d"},
		{"b", "c", "d", "a"},
		{"b", "d", "a", "c"},
		{"b", "d", "c", "a"},
		{"c", "a", "b", "d"},
		{"c", "a", "d", "b"},
		{"c", "b", "a", "d"},
		{"c", "b", "d", "a"},
		{"c", "d", "a", "b"},
		{"c", "d", "b", "a"},
		{"d", "a", "b", "c"},
		{"d", "a", "c", "b"},
		{"d", "b", "a", "c"},
		{"d", "b", "c", "a"},
		{"d", "c", "a", "b"},
		{"d", "c", "b", "a"},
	}

	for i := 0; i < len(perms)-2; i++ {
		if ok := perm(perms[i]); !ok {
			t.Errorf("[%d] got non-ok return", i)
		}
		if got, want := perms[i], perms[i+1]; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %v, want %v", i, got, want)
		}
	}

	if perm(perms[len(perms)-1]) {
		t.Errorf("got ok return value for last perm")
	}
}

func TestRev(t *testing.T) {
	for i, tc := range []struct {
		have, want []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"e", "d", "c", "b", "a"}},
	} {
		got := make([]string, len(tc.have))
		copy(got, tc.have)
		rev(got)
		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("[%d] got %v, want %v", i, got, tc.want)
		}
	}
}
