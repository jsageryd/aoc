package main

import "testing"

func TestSumOfNumbers(t *testing.T) {
	for n, tc := range []struct {
		input []byte
		sum   int
	}{
		{[]byte(`[1,2,3]`), 6},
		{[]byte(`{"a":2,"b":4}`), 6},
		{[]byte(`[[[3]]]`), 3},
		{[]byte(`{"a":{"b":4},"c":-1}`), 3},
		{[]byte(`{"a":[-1,1]}`), 0},
		{[]byte(`[-1,{"a":1}]`), 0},
		{[]byte(`[]`), 0},
		{[]byte(`{}`), 0},
	} {
		if got, want := sumOfNumbers(tc.input), tc.sum; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
