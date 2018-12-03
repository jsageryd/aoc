package main

import (
	"fmt"
)

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scan(&n); err != nil {
			break
		}
		input = append(input, n)
	}

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)
	_ = inputCopy

	cycles, loopCount := redistributionCyclesAndLoopSize(input)

	fmt.Printf("Part 1: %d\n", cycles)
	fmt.Printf("Part 2: %d\n", loopCount)
}

func redistributionCyclesAndLoopSize(memory []int) (cycles, loopSize int) {
	seen := [][]int{}

	for cycles = 1; ; cycles++ {
		distribute(memory, maxIdx(memory))
		for n := range seen {
			if intSlicesEqual(memory, seen[n]) {
				return cycles, cycles - n - 1
			}
		}
		memCpy := make([]int, len(memory))
		copy(memCpy, memory)
		seen = append(seen, memCpy)
	}
}

func distribute(s []int, idx int) {
	var blocks int
	blocks, s[idx] = s[idx], 0
	for ; blocks > 0; blocks-- {
		idx++
		s[idx%len(s)]++
	}
}

func intSlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for n := range s1 {
		if s1[n] != s2[n] {
			return false
		}
	}
	return true
}

func maxIdx(s []int) int {
	idx := 0
	for n := range s {
		if s[n] > s[idx] {
			idx = n
		}
	}
	return idx
}
