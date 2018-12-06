package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", len(reduce(input)))
	fmt.Printf("Part 2: %d\n", len(reduceAlternate(input)))
}

func reduce(polymer string) string {
	var stack []byte
	for _, b := range []byte(polymer) {
		if len(stack) > 0 && b^32 == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}

func reduceAlternate(polymer string) string {
	polymer = reduce(polymer)

	m := map[byte]struct{}{}

	for _, b := range []byte(polymer) {
		m[b|32] = struct{}{}
	}

	shortest := []byte(polymer)

	for unit := range m {
		var stack []byte
		for _, b := range []byte(polymer) {
			switch {
			case b|32 == unit:
				continue
			case len(stack) > 0 && b^32 == stack[len(stack)-1]:
				stack = stack[:len(stack)-1]
			default:
				stack = append(stack, b)
			}
		}
		if len(stack) < len(shortest) {
			shortest = stack
		}
	}

	return string(shortest)
}
