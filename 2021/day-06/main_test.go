package main

import (
	"fmt"
	"testing"
)

func TestSpawn(t *testing.T) {
	for n, tc := range []struct {
		input []int
		days  int
		want  []int
	}{
		{input: []int{3, 4, 3, 1, 2}, days: 1, want: []int{2, 3, 2, 0, 1}},
		{input: []int{3, 4, 3, 1, 2}, days: 2, want: []int{1, 2, 1, 6, 0, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 3, want: []int{0, 1, 0, 5, 6, 7, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 4, want: []int{6, 0, 6, 4, 5, 6, 7, 8, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 5, want: []int{5, 6, 5, 3, 4, 5, 6, 7, 7, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 6, want: []int{4, 5, 4, 2, 3, 4, 5, 6, 6, 7}},
		{input: []int{3, 4, 3, 1, 2}, days: 7, want: []int{3, 4, 3, 1, 2, 3, 4, 5, 5, 6}},
		{input: []int{3, 4, 3, 1, 2}, days: 8, want: []int{2, 3, 2, 0, 1, 2, 3, 4, 4, 5}},
		{input: []int{3, 4, 3, 1, 2}, days: 9, want: []int{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 10, want: []int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 11, want: []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 12, want: []int{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 13, want: []int{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 14, want: []int{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 15, want: []int{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7}},
		{input: []int{3, 4, 3, 1, 2}, days: 16, want: []int{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 17, want: []int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8}},
		{input: []int{3, 4, 3, 1, 2}, days: 18, want: []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}},
	} {
		if got, want := spawn(tc.input, tc.days), tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %d fish %v, want %d fish %v", n, len(got), got, len(want), want)
		}
	}
}
