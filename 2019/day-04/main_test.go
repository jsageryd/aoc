package main

import "testing"

func TestValidPassword(t *testing.T) {
	for n, tc := range []struct {
		pass  int
		valid bool
	}{
		{12345, false},   // 5 digits
		{1234567, false}, // 7 digits

		{122345, true}, // double digit and increasing
		{111123, true}, // double digit and increasing

		{111111, true},  // double digit
		{223450, false}, // decreasing
		{123789, false}, // no double digit
	} {
		if got, want := validPassword(tc.pass), tc.valid; got != want {
			t.Errorf("[%d] validPassword(%d) = %t, want %t", n, tc.pass, got, want)
		}
	}
}
