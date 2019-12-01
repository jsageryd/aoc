package main

import "testing"

func TestFuel(t *testing.T) {
	for n, tc := range []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	} {
		if got, want := fuel(tc.mass), tc.fuel; got != want {
			t.Errorf("[%d] fuel(%d) = %d, want %d", n, tc.mass, got, want)
		}
	}
}

func TestFuelWithFuel(t *testing.T) {
	for n, tc := range []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	} {
		if got, want := fuelWithFuel(tc.mass), tc.fuel; got != want {
			t.Errorf("[%d] fuelWithFuel(%d) = %d, want %d", n, tc.mass, got, want)
		}
	}
}
