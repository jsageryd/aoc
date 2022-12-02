package main

import "testing"

func TestTotalScore(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	if got, want := totalScore(input), 15; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestTotalScore2(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	if got, want := totalScore2(input), 12; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPlay(t *testing.T) {
	for n, tc := range []struct {
		ours, theirs string
		score        int
	}{
		{"A", "A", 1 + 3}, // rock + rock => draw; 1 + 3 = 4
		{"A", "B", 1 + 0}, // rock + paper => loss; 1 + 0 = 1
		{"A", "C", 1 + 6}, // rock + scissors => win; 1 + 6 = 7

		{"B", "A", 2 + 6}, // paper + rock => win; 2 + 6 = 8
		{"B", "B", 2 + 3}, // paper + paper => draw; 2 + 3 = 5
		{"B", "C", 2 + 0}, // paper + scissors => loss; 2 + 0 = 2

		{"C", "A", 3 + 0}, // scissors + rock => loss; 3 + 0 = 3
		{"C", "B", 3 + 6}, // scissors + paper => win; 3 + 6 = 9
		{"C", "C", 3 + 3}, // scissors + scissors => draw; 3 + 3 = 6
	} {
		if got, want := play(tc.ours, tc.theirs), tc.score; got != want {
			t.Errorf("[%d] play(%q, %q) = %d, want %d", n, tc.ours, tc.theirs, got, want)
		}
	}
}
