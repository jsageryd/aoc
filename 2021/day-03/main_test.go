package main

import "testing"

func Test(t *testing.T) {
	for _, tc := range []struct {
		desc string
		f    func(input []int) int
		want int
	}{
		{desc: "gammaRate", f: gammaRate, want: 0b10110},
		{desc: "epsilonRate", f: epsilonRate, want: 0b01001},
		{desc: "oxygenGeneratorRating", f: oxygenGeneratorRating, want: 0b10111},
		{desc: "co2ScrubberRating", f: co2ScrubberRating, want: 0b01010},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			input := []int{
				0b00100,
				0b11110,
				0b10110,
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11100,
				0b10000,
				0b11001,
				0b00010,
				0b01010,
			}

			if got, want := tc.f(input), tc.want; got != want {
				t.Errorf("got %[1]b (%[1]d), want %[2]b (%[2]d)", got, want)
			}
		})
	}
}
