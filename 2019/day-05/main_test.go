package main

import "testing"

func TestOpcode(t *testing.T) {
	for n, tc := range []struct {
		val  int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 10},
		{11, 11},
		{100, 0},
		{101, 1},
		{110, 10},
		{1000, 0},
		{1001, 1},
		{1010, 10},
		{10000, 0},
		{10001, 1},
		{10010, 10},
	} {
		if got, want := opcode(tc.val), tc.want; got != want {
			t.Errorf("[%d] opcode(%d) = %d, want %d", n, tc.val, got, want)
		}
	}
}

func TestParamMode(t *testing.T) {
	for n, tc := range []struct {
		val   int
		param int
		want  int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{100, 0, 1},
		{100, 1, 0},
		{1000, 0, 0},
		{1000, 1, 1},
	} {
		if got, want := paramMode(tc.val, tc.param), tc.want; got != want {
			t.Errorf("[%d] paramMode(%d, %d) = %d, want %d", n, tc.val, tc.param, got, want)
		}
	}
}
