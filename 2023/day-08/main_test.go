package main

import "testing"

func TestPart1(t *testing.T) {
	for n, tc := range []struct {
		input []string
		want  int
	}{
		{
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			want: 2,
		},
		{
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			want: 6,
		},
	} {
		if got, want := part1(tc.input), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
