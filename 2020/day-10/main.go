package main

import (
	"fmt"
	"sort"
)

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scanln(&n); err != nil {
			break
		}
		input = append(input, n)
	}

	one, three := joltageDiffFreq(input)

	fmt.Printf("Part 1: %d\n", one*three)
}

func joltageDiffFreq(joltages []int) (one, three int) {
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages)-1]+3) // device
	var lastN int
	for _, n := range joltages {
		switch n - lastN {
		case 1:
			one++
		case 3:
			three++
		}
		lastN = n
	}
	return one, three
}
