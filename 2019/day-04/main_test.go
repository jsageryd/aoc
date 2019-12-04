package main

import "testing"

func TestValidPassword(t *testing.T) {
	for n, tc := range []struct {
		pass         int
		strictDouble bool
		valid        bool
	}{
		// Part 1
		{12345, false, false},   // 5 digits
		{1234567, false, false}, // 7 digits

		{122345, false, true}, // double digit and increasing
		{111123, false, true}, // double digit and increasing

		{111111, false, true},  // double digit
		{223450, false, false}, // decreasing
		{123789, false, false}, // no double digit

		{123444, false, true}, // triple digits

		// Part 2
		{112233, true, true},  // strict double digit and increasing
		{123444, true, false}, // no strict double digit
		{111122, true, true},  // strict double digit and increasing
		{123455, true, true},  // strict double digit first and increasing
	} {
		if got, want := validPassword(tc.pass, tc.strictDouble), tc.valid; got != want {
			t.Errorf("[%d] validPassword(%d, %t) = %t, want %t", n, tc.pass, tc.strictDouble, got, want)
		}
	}
}

func BenchmarkNumberOfValidPasswords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numberOfValidPasswords(137683, 596253, false)
	}
}
