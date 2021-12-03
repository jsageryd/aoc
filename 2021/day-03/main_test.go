package main

import "testing"

func TestGammaRate(t *testing.T) {
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

	if got, want := gammaRate(input), 0b10110; got != want {
		t.Errorf("got %[1]b (%[1]d), want %[2]b (%[2]d)", got, want)
	}
}

func TestEpsilonRate(t *testing.T) {
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

	if got, want := epsilonRate(input), 0b01001; got != want {
		t.Errorf("got %[1]b (%[1]d), want %[2]b (%[2]d)", got, want)
	}
}
