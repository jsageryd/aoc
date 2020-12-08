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

	acc, ok := run(instructions)

	if got, want := acc, 5; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	if got, want := ok, false; got != want {
		t.Errorf("ok = %t, want %t", got, want)
	}
}

func TestFixAndRun(t *testing.T) {
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

	acc := fixAndRun(instructions)

	if got, want := acc, 8; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
