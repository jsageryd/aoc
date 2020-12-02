package main

import (
	"fmt"
	"testing"
)

func TestFindSum(t *testing.T) {
	for n, tc := range []struct {
		list     []int
		entries  int
		wantSum  int
		wantProd int
	}{
		{
			list:     []int{1721, 979, 366, 299, 675, 1456},
			entries:  2,
			wantSum:  2020,
			wantProd: 1721 * 299,
		},
		{
			list:     []int{1721, 979, 366, 299, 675, 1456},
			entries:  3,
			wantSum:  2020,
			wantProd: 979 * 366 * 675,
		},
	} {
		if got, want := findSum(tc.list, tc.entries, tc.wantSum), tc.wantProd; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestCombinations(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	k := 3

	wantCombs := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 5},
		{1, 2, 6},
		{1, 3, 4},
		{1, 3, 5},
		{1, 3, 6},
		{1, 4, 5},
		{1, 4, 6},
		{1, 5, 6},
		{2, 3, 4},
		{2, 3, 5},
		{2, 3, 6},
		{2, 4, 5},
		{2, 4, 6},
		{2, 5, 6},
		{3, 4, 5},
		{3, 4, 6},
		{3, 5, 6},
		{4, 5, 6},
	}

	var gotCombs [][]int

	combinations(s, k, func(comb []int) bool {
		c := make([]int, k)
		copy(c, comb)
		gotCombs = append(gotCombs, c)
		return true
	})

	if fmt.Sprint(gotCombs) != fmt.Sprint(wantCombs) {
		t.Errorf("got:\n%v\nwant:\n%v", gotCombs, wantCombs)
	}
}
