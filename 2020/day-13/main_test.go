package main

import (
	"strings"
	"testing"
)

const input = `939
7,13,x,x,59,x,31,19
`

func TestEarliestBus(t *testing.T) {
	estimate, schedule := parseInput(strings.NewReader(input))

	busID, waitTime := earliestBus(estimate, schedule)

	if got, want := busID, 59; got != want {
		t.Errorf("bus ID is %d, want %d", got, want)
	}

	if got, want := waitTime, 5; got != want {
		t.Errorf("wait time is %d, want %d", got, want)
	}
}

func TestEarliestTimestampForSubsequentDepartures(t *testing.T) {
	_, schedule := parseInput(strings.NewReader(input))

	earliest := earliestTimestampForSubsequentDepartures(schedule)

	if got, want := earliest, 1068781; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
