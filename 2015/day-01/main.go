package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", floor(input))
	fmt.Printf("Part 2: %d\n", firstBasementPosition(input))
}

func floor(input []byte) int {
	floor := 0

	for _, b := range input {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func firstBasementPosition(input []byte) int {
	floor := 0

	for n, b := range input {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			return n + 1
		}
	}

	return 0
}
