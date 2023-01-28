package main

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []string{
		"Sue 1: cars: 9, akitas: 3, goldfish: 0",
		"Sue 2: akitas: 9, children: 3, samoyeds: 9",
		"Sue 3: trees: 6, cars: 6, children: 4",
	}

	gotSues := parseInput(input)

	wantSues := []sue{
		{id: 1, things: map[string]int{"cars": 9, "akitas": 3, "goldfish": 0}},
		{id: 2, things: map[string]int{"akitas": 9, "children": 3, "samoyeds": 9}},
		{id: 3, things: map[string]int{"trees": 6, "cars": 6, "children": 4}},
	}

	if fmt.Sprint(gotSues) != fmt.Sprint(wantSues) {
		t.Errorf("got %v, want %v", gotSues, wantSues)
	}
}
