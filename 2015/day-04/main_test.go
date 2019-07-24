package main

import "testing"

func TestLeadingHexZeros(t *testing.T) {
	for n, tc := range []struct {
		input []byte
		want  int
	}{
		{[]byte{0x10}, 0},
		{[]byte{0x01}, 1},
		{[]byte{0x00}, 2},
		{[]byte{0x00, 0x01}, 3},
		{[]byte{0x00, 0x00}, 4},
		{[]byte{0x00, 0x00, 0x01}, 5},
		{[]byte{0x00, 0x00, 0x00}, 6},
	} {
		if got, want := leadingHexZeros(tc.input), tc.want; got != want {
			t.Errorf("[%d] leadingHexZeros(%#x) = %d, want %d", n, tc.input, got, want)
		}
	}
}
