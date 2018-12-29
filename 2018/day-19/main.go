package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var reg [6]int

	run(&reg, input)

	fmt.Printf("Part 1: %d\n", reg[0])
}

type Inst struct {
	op      func(reg *[6]int, a, b, c int)
	a, b, c int
}

func run(reg *[6]int, programStrs []string) {
	var ipReg int
	fmt.Sscanf(programStrs[0], "#ip %d", &ipReg)
	programStrs = programStrs[1:]

	program := make([]Inst, 0, len(programStrs))

	for n := range programStrs {
		var op string
		var inst Inst
		fmt.Sscanf(programStrs[n], "%s %d %d %d", &op, &inst.a, &inst.b, &inst.c)
		inst.op = ops[op]
		program = append(program, inst)
	}

	for ip := 0; ip >= 0 && ip < len(program); ip++ {
		inst := program[ip]
		reg[ipReg] = ip
		inst.op(reg, inst.a, inst.b, inst.c)
		ip = reg[ipReg]
	}
}

var ops = map[string]func(reg *[6]int, a, b, c int){
	"addr": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] + reg[b] },
	"addi": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] + b },
	"mulr": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] * reg[b] },
	"muli": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] * b },
	"banr": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] & reg[b] },
	"bani": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] & b },
	"borr": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] | reg[b] },
	"bori": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] | b },
	"setr": func(reg *[6]int, a, b, c int) { reg[c] = reg[a] },
	"seti": func(reg *[6]int, a, b, c int) { reg[c] = a },
	"gtir": func(reg *[6]int, a, b, c int) {
		if a > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"gtri": func(reg *[6]int, a, b, c int) {
		if reg[a] > b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"gtrr": func(reg *[6]int, a, b, c int) {
		if reg[a] > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqir": func(reg *[6]int, a, b, c int) {
		if a == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqri": func(reg *[6]int, a, b, c int) {
		if reg[a] == b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
	"eqrr": func(reg *[6]int, a, b, c int) {
		if reg[a] == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	},
}
