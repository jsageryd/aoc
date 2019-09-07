package main

import (
	"fmt"
	"testing"
)

func TestCircuit_MeasureAll(t *testing.T) {
	for n, tc := range []struct {
		instructions []string
		signals      map[string]int
	}{
		{
			instructions: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			signals: map[string]int{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
	} {
		c := newCircuit(tc.instructions)

		if got, want := c.measureAll(), tc.signals; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}
