package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	for n, tc := range []struct {
		code string
		want []int
	}{
		{"", []int{}},
		{"0", []int{0}},
		{"0,1", []int{0, 1}},
		{"0,2,1", []int{0, 2, 1}},
	} {
		if got, want := parse(tc.code), tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] parse(%q) = %v, want %v", n, tc.code, got, want)
		}
	}
}

func TestRun(t *testing.T) {
	for n, tc := range []struct {
		code []int
		want []int
	}{
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		got, err := run(tc.code)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("[%d] run(%v) = %v, want %v", n, tc.code, got, tc.want)
		}
	}
}
