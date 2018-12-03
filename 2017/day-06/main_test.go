package main

import "testing"

func TestRedistributionCyclesAndLoopSize(t *testing.T) {
	for n, tc := range []struct {
		in       []int
		cycles   int
		loopSize int
	}{
		{[]int{0, 2, 7, 0}, 5, 4},
	} {
		cycles, loopSize := redistributionCyclesAndLoopSize(tc.in)

		if got, want := cycles, tc.cycles; got != want {
			t.Errorf("[%d] cycles = %d, want %d", n, got, want)
		}

		if got, want := loopSize, tc.loopSize; got != want {
			t.Errorf("[%d] loopSize = %d, want %d", n, got, want)
		}
	}
}

func TestDistribute(t *testing.T) {
	for n, tc := range []struct {
		in  []int
		idx int
		out []int
	}{
		{[]int{0, 1, 2, 3}, 0, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, 1, []int{0, 0, 3, 3}},
		{[]int{0, 1, 2, 3}, 2, []int{1, 1, 0, 4}},
		{[]int{0, 1, 2, 3}, 3, []int{1, 2, 3, 0}},
		{[]int{3, 0, 1, 2}, 0, []int{0, 1, 2, 3}},
	} {
		distribute(tc.in, tc.idx)
		for i := range tc.in {
			if tc.in[i] != tc.out[i] {
				t.Errorf("[%d] got %v, want %v", n, tc.in, tc.out)
				break
			}
		}
	}
}

func TestMaxIdx(t *testing.T) {
	for n, tc := range []struct {
		in  []int
		out int
	}{
		{[]int{0}, 0},
		{[]int{0, 0}, 0},
		{[]int{1, 1}, 0},
		{[]int{0, 1, 2, 3}, 3},
		{[]int{1, 2, 3, 0}, 2},
		{[]int{2, 3, 0, 1}, 1},
		{[]int{3, 0, 1, 2}, 0},
	} {
		if got, want := maxIdx(tc.in), tc.out; got != want {
			t.Errorf("[%d] maxIdx(%v) = %d, want %d", n, tc.in, got, want)
		}
	}
}

func TestIntSlicesEqual(t *testing.T) {
	for n, tc := range []struct {
		s1    []int
		s2    []int
		equal bool
	}{
		{[]int{}, []int{}, true},
		{[]int{0}, []int{0}, true},
		{[]int{0, 1, 2}, []int{0, 1, 2}, true},
		{[]int{0}, []int{}, false},
		{[]int{}, []int{0}, false},
		{[]int{0, 1, 2}, []int{1, 2, 0}, false},
		{[]int{0, 1, 2}, []int{2, 0, 1}, false},
	} {
		if got, want := intSlicesEqual(tc.s1, tc.s2), tc.equal; got != want {
			t.Errorf("[%d] got %t, want %t", n, got, want)
		}
	}
}
