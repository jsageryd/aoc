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
	fmt.Printf("Part 2: %d\n", possibleArrangements(input))
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

func possibleArrangements(joltages []int) int {
	joltages = append(joltages, 0) // outlet
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages)-1]+3) // device

	type adapter struct {
		connections      []*adapter
		totalConnections int
	}

	adapters := make([]*adapter, len(joltages))
	for i := range adapters {
		adapters[i] = &adapter{}
	}

	for i := 0; i < len(joltages); i++ {
		for j := i + 1; j < len(joltages) && joltages[j]-joltages[i] <= 3; j++ {
			adapters[i].connections = append(adapters[i].connections, adapters[j])
		}
	}

	var rec func(a *adapter) int

	rec = func(a *adapter) int {
		if a.totalConnections != 0 {
			return a.totalConnections
		}
		if len(a.connections) == 0 {
			return 1
		}
		for _, c := range a.connections {
			a.totalConnections += rec(c)
		}
		return a.totalConnections
	}

	return rec(adapters[0])
}
