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

	fmt.Printf("Part 1: %d\n", run(input))
}

func run(instructions []string) (acc int) {
	seen := make(map[int]struct{})
	var cur int

	for {
		if _, ok := seen[cur]; ok {
			break
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

	return acc
}
