package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
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
