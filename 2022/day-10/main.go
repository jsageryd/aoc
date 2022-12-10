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

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var sum int

	c := newCPU(input)

	c.run(func() {
		if (c.cycle-20)%40 == 0 {
			sum += c.cycle * c.x
		}
	})

	return sum
}

type cpu struct {
	program  []string
	cycle    int
	cur      int
	curSpent int
	x        int
}

func newCPU(program []string) *cpu {
	return &cpu{
		program: program,
		x:       1,
	}
}

// run runs the program, calling f at the beginning of each cycle.
func (c *cpu) run(f func()) {
	cost := map[string]int{
		"addx": 2,
		"noop": 1,
	}

	for c.cur < len(c.program) {
		c.cycle++
		c.curSpent++

		f()

		inst := strings.Fields(c.program[c.cur])

		if c.curSpent < cost[inst[0]] {
			continue
		}

		switch inst[0] {
		case "addx":
			val, _ := strconv.Atoi(inst[1])
			c.x += val
		case "noop":
			// noop
		}

		c.cur++
		c.curSpent = 0
	}
}
