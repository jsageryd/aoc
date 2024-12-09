package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []byte) int {
	var sum int

	var pos int

	yield := func(fileID int) {
		sum += pos * fileID
		pos++
	}

	ints := parse(input)

	N := len(ints) - 1

	for n := 0; n <= N; n++ {
		if n%2 == 0 { // file
			fileID := n / 2

			for range ints[n] {
				yield(fileID)
			}
		} else { // space
			fileID := N / 2

			for ints[n] > 0 && ints[N] > 0 {
				yield(fileID)

				ints[n]--
				ints[N]--

				if ints[N] == 0 {
					N -= 2
					fileID = N / 2
				}
			}
		}
	}

	return sum
}

func parse(input []byte) []int {
	ints := make([]int, 0, len(input))

	for n := range bytes.TrimSpace(input) {
		ints = append(ints, int(input[n]-'0'))
	}

	return ints
}
