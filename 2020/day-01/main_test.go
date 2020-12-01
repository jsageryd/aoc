package main

import "testing"

func TestFindSumOfTwo(t *testing.T) {
	for n, tc := range []struct {
		list    []int
		wantSum int
		n1, n2  int
	}{
		{
			list:    []int{1721, 979, 366, 299, 675, 1456},
			wantSum: 2020,
			n1:      1721,
			n2:      299,
		},
	} {
		n1, n2 := findSumOfTwo(tc.list, tc.wantSum)

		if n1 != tc.n1 || n2 != tc.n2 {
			t.Errorf("[%d] got %d, %d, want %d, %d", n, n1, n2, tc.n1, tc.n2)
		}
	}
}

func TestFindSumOfThree(t *testing.T) {
	for n, tc := range []struct {
		list       []int
		wantSum    int
		n1, n2, n3 int
	}{
		{
			list:    []int{1721, 979, 366, 299, 675, 1456},
			wantSum: 2020,
			n1:      979,
			n2:      366,
			n3:      675,
		},
	} {
		n1, n2, n3 := findSumOfThree(tc.list, tc.wantSum)

		if n1 != tc.n1 || n2 != tc.n2 || n3 != tc.n3 {
			t.Errorf("[%d] got %d, %d, %d, want %d, %d, %d", n, n1, n2, n3, tc.n1, tc.n2, tc.n3)
		}
	}
}
