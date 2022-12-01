package main

import "testing"

func TestMaxCalories(t *testing.T) {
	input := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	first, second, third := maxCalories(input)

	if got, want := first, 24000; got != want {
		t.Errorf("first is %d, want %d", got, want)
	}

	if got, want := second, 11000; got != want {
		t.Errorf("second is %d, want %d", got, want)
	}

	if got, want := third, 10000; got != want {
		t.Errorf("third is %d, want %d", got, want)
	}
}
