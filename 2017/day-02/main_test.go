package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	for n, tc := range []struct {
		in  [][]int
		out int
	}{
		{[][]int{
			{5, 1, 9, 5},
			{7, 5, 3},
			{2, 4, 6, 8},
		},
			18,
		},
	} {
		if got, want := checksum(tc.in), tc.out; got != want {
			t.Errorf("[%d] checksum(%v) = %d, want %d", n, tc.in, got, want)
		}
	}
}

func TestChecksumDivisibles(t *testing.T) {
	for n, tc := range []struct {
		in  [][]int
		out int
	}{
		{[][]int{
			{5, 9, 2, 8},
			{9, 4, 7, 3},
			{3, 8, 6, 5},
		},
			9,
		},
	} {
		if got, want := checksumDivisibles(tc.in), tc.out; got != want {
			t.Errorf("[%d] checksumDivisibles(%v) = %d, want %d", n, tc.in, got, want)
		}
	}
}
