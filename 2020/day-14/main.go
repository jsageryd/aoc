package main

import (
	"bufio"
	"fmt"
	"os"
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
	applyMask := func(val int, mask string) int {
		for i := range mask {
			switch mask[len(mask)-i-1] {
			case '0':
				val &= ^(1 << i)
			case '1':
				val |= 1 << i
			}
		}
		return val
	}

	var sum int
	mem := make(map[int]int)

	var mask string

	for _, row := range input {
		switch {
		case strings.HasPrefix(row, "mask"):
			mask = strings.TrimPrefix(row, "mask = ")
		case strings.HasPrefix(row, "mem"):
			var addr, val int
			fmt.Sscanf(row, "mem[%d] = %d", &addr, &val)
			mem[addr] = applyMask(val, mask)
		}
	}

	for _, v := range mem {
		sum += v
	}

	return sum
}
