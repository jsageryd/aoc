package main

import "testing"

func TestRun(t *testing.T) {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	if got, want := run(instructions), 5; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
