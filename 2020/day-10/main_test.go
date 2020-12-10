package main

import "testing"

func TestJoltageDiffFreq(t *testing.T) {
	for n, tc := range []struct {
		joltages                    []int
		wantOne, wantTwo, wantThree int
	}{
		{
			joltages: []int{
				16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4,
			},
			wantOne:   7,
			wantThree: 5,
		},
		{
			joltages: []int{
				28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19,
				38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3,
			},
			wantOne:   22,
			wantThree: 10,
		},
	} {
		gotOne, gotThree := joltageDiffFreq(tc.joltages)

		if gotOne != tc.wantOne || gotThree != tc.wantThree {
			t.Errorf("[%d] got %d, %d, want %d, %d", n, gotOne, gotThree, tc.wantOne, tc.wantThree)
		}
	}
}
