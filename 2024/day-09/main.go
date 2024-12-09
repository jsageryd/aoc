package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func part2(input []byte) int {
	type file struct {
		id   int
		size int
	}

	fileMap := make(map[int][]file)

	ints := parse(input)
	origInts := slices.Clone(ints)

	for N := len(ints) - 1; N >= 0; N -= 2 {
		fileID := N / 2

		for n := 1; n < N; n += 2 {
			if ints[n] >= ints[N] {
				fileMap[n] = append(fileMap[n], file{fileID, ints[N]})
				ints[n] -= ints[N]
				ints[N] = 0
				break
			}
		}
	}

	for n := 0; n < len(ints); n += 2 {
		fileMap[n] = []file{{n / 2, ints[n]}}
	}

	var sum int

	var pos int

	for n := range ints {
		if files, ok := fileMap[n]; ok {
			for _, f := range files {
				for range f.size {
					sum += pos * f.id
					pos++
				}
			}
		}

		if n%2 != 0 {
			pos += ints[n]
		} else {
			if ints[n] == 0 {
				pos += origInts[n]
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
