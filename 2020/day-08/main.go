package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	acc, _ := run(input)
	fmt.Printf("Part 1: %d\n", acc)

	acc = fixAndRun(input)
	fmt.Printf("Part 2: %d\n", acc)
}

func run(instructions []string) (acc int, ok bool) {
	seen := make(map[int]struct{})
	var cur int

	for cur < len(instructions) {
		if _, ok := seen[cur]; ok {
			return acc, false
		}
		seen[cur] = struct{}{}
		inst := instructions[cur]
		split := strings.Split(inst, " ")
		op, argStr := split[0], split[1]
		arg, _ := strconv.Atoi(argStr)
		switch op {
		case "acc":
			acc += arg
		case "jmp":
			cur += arg
			continue
		}
		cur++
	}

	return acc, true
}

func fixAndRun(instructions []string) (acc int) {
	replace := func(idx int) bool {
		split := strings.Split(instructions[idx], " ")
		op, argStr := split[0], split[1]
		switch op {
		case "jmp":
			op = "nop"
		case "nop":
			op = "jmp"
		default:
			return false
		}
		instructions[idx] = op + " " + argStr
		return true
	}

	for cur := range instructions {
		if replace(cur) {
			if acc, ok := run(instructions); ok {
				return acc
			} else {
				replace(cur)
			}
		}
	}

	return 0
}
