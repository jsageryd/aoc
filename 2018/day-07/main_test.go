package main

import "testing"

var input = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func TestOrderedSteps(t *testing.T) {
	if got, want := orderedSteps(parseInput(input)), "CABDFE"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTimeToCompleteSteps(t *testing.T) {
	if got, want := timeToCompleteSteps(parseInput(input), 2, 0), 15; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCost(t *testing.T) {
	for n, tc := range []struct {
		step string
		cost int
	}{
		{"A", 61},
		{"B", 62},
		{"C", 63},
		{"D", 64},
		{"E", 65},
		{"F", 66},
		{"G", 67},
		{"H", 68},
		{"I", 69},
		{"J", 70},
		{"K", 71},
		{"L", 72},
		{"M", 73},
		{"N", 74},
		{"O", 75},
		{"P", 76},
		{"Q", 77},
		{"R", 78},
		{"S", 79},
		{"T", 80},
		{"U", 81},
		{"V", 82},
		{"W", 83},
		{"X", 84},
		{"Y", 85},
		{"Z", 86},
	} {
		if got, want := cost(tc.step, 60), tc.cost; got != want {
			t.Errorf("[%d] cost(%q) = %d, want %d", n, tc.step, got, want)
		}
	}
}
