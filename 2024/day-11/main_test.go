package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "125 17"

	if got, want := part1(input), 55312; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{999, 3},
		{1000, 4},
		{9999, 4},
		{10000, 5},
	} {
		if got, want := digits(tc.n), tc.want; got != want {
			t.Errorf("digits(%d) = %d, want %d", tc.n, got, want)
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		n           int
		left, right int
	}{
		{0, 0, 0},
		{1, 1, 0},
		{9, 9, 0},
		{10, 1, 0},
		{99, 9, 9},
		{100, 10, 0},
		{999, 99, 9},
		{1000, 10, 0},
		{1234, 12, 34},
		{9999, 99, 99},
		{10000, 100, 0},
		{100000, 100, 0},
		{123456, 123, 456},
	} {
		gotLeft, gotRight := split(tc.n)

		if gotLeft != tc.left || gotRight != tc.right {
			t.Errorf("split(%d) = %d, %d, want %d, %d", tc.n, gotLeft, gotRight, tc.left, tc.right)
		}
	}
}
