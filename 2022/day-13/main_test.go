package main

import (
	"encoding/json"
	"testing"
)

var input = []byte(`[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`)

func TestPart1(t *testing.T) {
	if got, want := part1(input), 13; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 140; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCompare(t *testing.T) {
	for n, tc := range []struct {
		a    string
		b    string
		want int
	}{
		{
			"[]",
			"[]",
			0,
		},
		{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			-1,
		},
		{
			"[[1],[2,3,4]]",
			"[[1],4]",
			-1,
		},
		{
			"[9]",
			"[[8,7,6]]",
			1,
		},
		{
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			-1,
		},
		{
			"[7,7,7,7]",
			"[7,7,7]",
			1,
		},
		{
			"[]",
			"[3]",
			-1,
		},
		{
			"[[[]]]",
			"[[]]",
			1,
		},
		{
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
			1,
		},
	} {
		var a, b any

		if err := json.Unmarshal([]byte(tc.a), &a); err != nil {
			t.Fatalf("[%d] unexpected error: %v", n, err)
		}

		if err := json.Unmarshal([]byte(tc.b), &b); err != nil {
			t.Fatalf("[%d] unexpected error: %v", n, err)
		}

		if got, want := compare(a, b), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
