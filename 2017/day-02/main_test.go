package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	for n, tc := range []struct {
		in  []byte
		out int
	}{
		{[]byte("5	1	9	5\n7	5	3\n2	4	6	8\n"), 18},
	} {
		if got, want := checksum(tc.in), tc.out; got != want {
			t.Errorf("[%d] checksum(%q) = %d, want %d", n, tc.in, got, want)
		}
	}
}

func TestChecksumDivisibles(t *testing.T) {
	for n, tc := range []struct {
		in  []byte
		out int
	}{
		{[]byte("5	9	2	8\n9	4	7	3\n3	8	6	5\n"), 9},
	} {
		if got, want := checksumDivisibles(tc.in), tc.out; got != want {
			t.Errorf("[%d] checksumDivisibles(%q) = %d, want %d", n, tc.in, got, want)
		}
	}
}
