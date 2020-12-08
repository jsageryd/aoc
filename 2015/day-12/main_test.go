package main

import "testing"

func TestSumOfNumbers(t *testing.T) {
	for n, tc := range []struct {
		input          []byte
		skipRedObjects bool
		sum            int
	}{
		// Part 1
		{[]byte(`[1,2,3]`), false, 6},
		{[]byte(`{"a":2,"b":4}`), false, 6},
		{[]byte(`[[[3]]]`), false, 3},
		{[]byte(`{"a":{"b":4},"c":-1}`), false, 3},
		{[]byte(`{"a":[-1,1]}`), false, 0},
		{[]byte(`[-1,{"a":1}]`), false, 0},
		{[]byte(`[]`), false, 0},
		{[]byte(`{}`), false, 0},

		// Part 2
		{[]byte(`[1,2,3]`), true, 6},
		{[]byte(`[1,{"c":"red","b":2},3]`), true, 4},
		{[]byte(`{"d":"red","e":[1,2,3,4],"f":5}`), true, 0},
		{[]byte(`[1,"red",5]`), true, 6},
	} {
		if got, want := sumOfNumbers(tc.input, tc.skipRedObjects), tc.sum; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
