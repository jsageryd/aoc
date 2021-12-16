package main

import (
	"fmt"
	"testing"
)

var distances = map[string]map[string]int{
	"London": {
		"Belfast": 518,
		"Dublin":  464,
	},
	"Dublin": {
		"Belfast": 141,
		"London":  464,
	},
	"Belfast": {
		"Dublin": 141,
		"London": 518,
	},
}

func TestFindShortestRoute(t *testing.T) {
	gotRoute, gotDistance := findShortestRoute(distances)

	wantRoute := []string{"London", "Dublin", "Belfast"}
	wantDistance := 605

	if fmt.Sprint(gotRoute) != fmt.Sprint(wantRoute) || gotDistance != wantDistance {
		t.Errorf("got %v (%d), want %v (%d)", gotRoute, gotDistance, wantRoute, wantDistance)
	}
}

func TestFindLongestRoute(t *testing.T) {
	gotRoute, gotDistance := findLongestRoute(distances)

	wantRoute := []string{"Belfast", "London", "Dublin"}
	wantDistance := 982

	if fmt.Sprint(gotRoute) != fmt.Sprint(wantRoute) || gotDistance != wantDistance {
		t.Errorf("got %v (%d), want %v (%d)", gotRoute, gotDistance, wantRoute, wantDistance)
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
