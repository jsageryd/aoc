package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var count int

	ranges, ids := parse(input)

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}

	return count
}

func part2(input []string) int {
	var count int

	ranges, _ := parse(input)

	for _, r := range linearize(ranges) {
		count += r[1] - r[0] + 1
	}

	return count
}

func linearize(ranges [][2]int) [][2]int {
	var done bool

	for !done {
		done = true

		for n, r1 := range ranges {
			for _, r2 := range ranges {
				// If the ranges don't overlap, skip
				// [r1   ] [r2   ]
				// [r2   ] [r1   ]
				if r1[1] < r2[0]-1 || r2[1] < r1[0]-1 {
					continue
				}

				// Merge the overlapping ranges
				r1[0] = min(r1[0], r2[0])
				r1[1] = max(r1[1], r2[1])

				if ranges[n] != r1 {
					done = false
				}

				ranges[n] = r1
			}
		}
	}

	slices.SortFunc(ranges, func(a, b [2]int) int {
		return slices.Compare(a[:], b[:])
	})

	return slices.Compact(ranges)
}

func parse(input []string) (ranges [][2]int, ids []int) {
	var readingIDs bool

	for _, line := range input {
		if len(line) == 0 {
			readingIDs = true
			continue
		}

		if readingIDs {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		} else {
			var start, end int
			fmt.Sscanf(line, "%d-%d", &start, &end)
			ranges = append(ranges, [2]int{start, end})
		}
	}

	return ranges, ids
}
