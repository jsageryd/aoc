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

	if got, want := maxCalories(input), 24000; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
