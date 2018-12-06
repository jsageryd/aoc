package main

import "testing"

func TestProcess(t *testing.T) {
	input := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}

	registers, maxAlloc := process(input)

	wantRegisters := map[string]int{
		"a": 1,
		"c": -10,
	}

	if len(registers) != len(wantRegisters) {
		t.Fatalf("got %v, want %v", registers, wantRegisters)
	}

	for k := range wantRegisters {
		if wantRegisters[k] != registers[k] {
			t.Fatalf("got %v, want %v", registers, wantRegisters)
		}
	}

	if got, want := maxAlloc, 10; got != want {
		t.Errorf("maxAlloc = %d, want %d", got, want)
	}
}

func TestLargestRegisterValue(t *testing.T) {
	for n, tc := range []struct {
		regs    map[string]int
		largest int
	}{
		{map[string]int{}, 0},
		{map[string]int{"foo": 1}, 1},
		{map[string]int{"foo": 1, "bar": 2}, 2},
		{map[string]int{"foo": 0, "bar": -1}, 0},
	} {
		if got, want := largestRegisterValue(tc.regs), tc.largest; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
