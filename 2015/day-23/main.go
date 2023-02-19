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
	c := newCPU(input)
	c.run()
	return c.reg["b"]
}

type cpu struct {
	program []string
	cur     int
	reg     map[string]int
}

func newCPU(program []string) *cpu {
	return &cpu{
		program: program,
		reg:     make(map[string]int),
	}
}

func (c *cpu) run() {
	for c.cur < len(c.program) {
		inst := strings.Fields(c.program[c.cur])

		switch inst[0] {
		case "hlf":
			c.reg[inst[1]] /= 2
		case "tpl":
			c.reg[inst[1]] *= 3
		case "inc":
			c.reg[inst[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(inst[1])
			c.cur += offset
			continue
		case "jie":
			r := inst[1][:len(inst[1])-1]
			if c.reg[r]%2 == 0 {
				offset, _ := strconv.Atoi(inst[2])
				c.cur += offset
				continue
			}
		case "jio":
			r := inst[1][:len(inst[1])-1]
			if c.reg[r] == 1 {
				offset, _ := strconv.Atoi(inst[2])
				c.cur += offset
				continue
			}
		}

		c.cur++
	}
}
