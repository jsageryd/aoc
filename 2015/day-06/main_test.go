package main

import "testing"

func TestApplyInstruction(t *testing.T) {
	for n, tc := range []struct {
		inst   string
		before grid
		after  grid
	}{
		{
			inst: "turn on 0,0 through 2,0",
			before: grid{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{1, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "turn off 0,0 through 2,0",
			before: grid{
				[]int{1, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "toggle 0,0 through 2,0",
			before: grid{
				[]int{0, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{1, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "toggle 0,0 through 2,2",
			before: grid{
				[]int{0, 1, 1},
				[]int{0, 1, 0},
				[]int{0, 0, 1},
			},
			after: grid{
				[]int{1, 0, 0},
				[]int{1, 0, 1},
				[]int{1, 1, 0},
			},
		},
	} {
		if got, want := applyInstruction(tc.inst, tc.before), tc.after; got.String() != want.String() {
			t.Errorf("[%d]\n%s\n\nbefore:\n%s\n\nafter:\n%s\n\nwant:\n%s", n, tc.inst, tc.before, got, want)
		}
	}
}

func TestApplyInstruction2(t *testing.T) {
	for n, tc := range []struct {
		inst   string
		before grid
		after  grid
	}{
		{
			inst: "turn on 0,0 through 2,0",
			before: grid{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{1, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "turn off 0,0 through 2,0",
			before: grid{
				[]int{1, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "toggle 0,0 through 2,0",
			before: grid{
				[]int{0, 1, 1},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			after: grid{
				[]int{2, 3, 3},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			inst: "toggle 0,0 through 2,2",
			before: grid{
				[]int{0, 1, 1},
				[]int{0, 1, 0},
				[]int{0, 0, 1},
			},
			after: grid{
				[]int{2, 3, 3},
				[]int{2, 3, 2},
				[]int{2, 2, 3},
			},
		},
	} {
		after := applyInstruction2(tc.inst, tc.before)

		if got, want := after, tc.after; got.String() != want.String() {
			t.Errorf("[%d]\n%s\n\nbefore:\n%s\n\nafter:\n%s\n\nwant:\n%s", n, tc.inst, tc.before, got, want)
		}

		if got, want := totalBrightness(after), totalBrightness(tc.after); got != want {
			t.Errorf("[%d] total brightness is %d, want %d\n\n", n, got, want)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out instruction
	}{
		{"turn on 0,0 through 999,999", instruction{actionOn, coord{0, 0}, coord{999, 999}}},
		{"toggle 0,0 through 999,0", instruction{actionToggle, coord{0, 0}, coord{999, 0}}},
		{"turn off 499,499 through 500,500", instruction{actionOff, coord{499, 499}, coord{500, 500}}},
	} {
		if got, want := parseInstruction(tc.in), tc.out; got != want {
			t.Errorf("[%d] parseInstruction(%q) = %+v, want %+v", n, tc.in, got, want)
		}
	}
}

func TestParseCoord(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out coord
	}{
		{"0,0", coord{0, 0}},
		{"123,456", coord{123, 456}},
	} {
		if got, want := parseCoord(tc.in), tc.out; got != want {
			t.Errorf("[%d] parseCoord(%q) = %+v, want %+v", n, tc.in, got, want)
		}
	}
}
