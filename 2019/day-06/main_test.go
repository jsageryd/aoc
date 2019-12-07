package main

import (
	"testing"
)

func TestTotalOrbits(t *testing.T) {
	for n, tc := range []struct {
		object *Object
		want   int
	}{
		{parse([]string{"COM)A"}), 1},
		{parse([]string{"COM)A", "A)B"}), 3},
		{parse([]string{"COM)A", "A)B", "A)C"}), 5},
		{parse([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}), 42},
		{parse([]string{"COM)B", "C)D", "B)C", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}), 42},
	} {
		if got, want := totalOrbits(tc.object), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestParse(t *testing.T) {
	for n, tc := range []struct {
		input []string
		want  string
	}{
		{
			[]string{"COM)A"},
			"COM(A)",
		},
		{
			[]string{"COM)A", "A)B"},
			"COM(A(B))",
		},
		{
			[]string{"COM)A", "A)B", "A)C"},
			"COM(A(B,C))",
		},
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			"COM(B(C(D(E(F,J(K(L))),I)),G(H)))",
		},
	} {
		if got, want := parse(tc.input).String(), tc.want; got != want {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}
