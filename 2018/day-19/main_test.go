package main

import (
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	program := []string{
		"#ip 0",
		"seti 5 0 1",
		"seti 6 0 2",
		"addi 0 1 0",
		"addr 1 2 3",
		"setr 1 0 0",
		"seti 8 0 4",
		"seti 9 0 5",
	}

	wantReg := [6]int{6, 5, 6, 0, 0, 9}

	var reg [6]int

	run(&reg, program)

	if reg != wantReg {
		t.Errorf("got %v, want %v", reg, wantReg)
	}
}

func BenchmarkRun(b *testing.B) {
	input, _ := os.ReadFile("input")
	programStrs := strings.Split(string(input), "\n")

	for n := 0; n < b.N; n++ {
		var reg [6]int
		run(&reg, programStrs)
	}
}
