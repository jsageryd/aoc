package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	regs, maxAlloc := process(input)

	fmt.Printf("Part 1: %d\n", largestRegisterValue(regs))
	fmt.Printf("Part 2: %d\n", maxAlloc)
}

func largestRegisterValue(regs map[string]int) int {
	largest := 0
	for _, v := range regs {
		if v > largest {
			largest = v
		}
	}
	return largest
}

type instruction struct {
	reg1 string
	op1  string
	val1 int
	reg2 string
	op2  string
	val2 int
}

func process(instructions []string) (regs map[string]int, maxAlloc int) {
	regs = map[string]int{}

	op1f := map[string]func(string, int){
		"inc": func(reg string, n int) {
			regs[reg] += n
		},
		"dec": func(reg string, n int) {
			regs[reg] -= n
		},
	}

	for _, instStr := range instructions {
		var inst instruction

		fmt.Sscanf(instStr, "%s %s %d if %s %s %d", &inst.reg1, &inst.op1, &inst.val1, &inst.reg2, &inst.op2, &inst.val2)

		switch inst.op2 {
		case ">":
			if regs[inst.reg2] > inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		case "<":
			if regs[inst.reg2] < inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		case ">=":
			if regs[inst.reg2] >= inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		case "<=":
			if regs[inst.reg2] <= inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		case "==":
			if regs[inst.reg2] == inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		case "!=":
			if regs[inst.reg2] != inst.val2 {
				op1f[inst.op1](inst.reg1, inst.val1)
			}
		default:
			log.Fatalf("unknown operator: %q", inst.op2)
		}

		if regs[inst.reg1] > maxAlloc {
			maxAlloc = regs[inst.reg1]
		}
	}

	return regs, maxAlloc
}
