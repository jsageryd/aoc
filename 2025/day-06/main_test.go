package main

import (
	"slices"
	"testing"
)

var input = []string{
	"123 328  51 64",
	" 45 64  387 23",
	"  6 98  215 314",
	"*   +   *   +",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 4277556; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSlice(t *testing.T) {
	gotSlices := slice(input)

	wantSlices := [][]string{
		{"123", " 45", "  6", "*  "},
		{"328", "64 ", "98 ", "+  "},
		{" 51", "387", "215", "*  "},
		{"64 ", "23 ", "314", "+  "},
	}

	if !slices.EqualFunc(gotSlices, wantSlices, slices.Equal) {
		t.Errorf("got:\n%q\n\nwant:\n%q", gotSlices, wantSlices)
	}
}
