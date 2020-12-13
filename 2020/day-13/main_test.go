package main

import (
	"strings"
	"testing"
)

func TestEarliestBus(t *testing.T) {
	const input = `939
7,13,x,x,59,x,31,19
`

	estimate, buses := parseInput(strings.NewReader(input))

	busID, waitTime := earliestBus(estimate, buses)

	if got, want := busID, 59; got != want {
		t.Errorf("bus ID is %d, want %d", got, want)
	}

	if got, want := waitTime, 5; got != want {
		t.Errorf("wait time is %d, want %d", got, want)
	}
}
