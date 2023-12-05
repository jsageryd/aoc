package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []byte) int {
	seeds, maps := parse(input)

	var locations []int

	for _, cur := range seeds {
		for _, m := range maps {
			if next, ok := m[cur]; ok {
				cur = next
			}
		}
		locations = append(locations, cur)
	}

	return slices.Min(locations)
}

func parse(input []byte) (seeds []int, maps []map[int]int) {
	parts := strings.Split(string(input), "\n\n")

	for _, s := range strings.Fields(parts[0][7:]) {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}

	for _, mapStr := range parts[1:] {
		m := make(map[int]int)

		for _, line := range strings.Split(strings.TrimSpace(mapStr), "\n")[1:] {
			var dst, src, length int

			fmt.Sscanf(line, "%d %d %d", &dst, &src, &length)

			for n := 0; n < length; n++ {
				m[src+n] = dst + n
			}
		}

		maps = append(maps, m)
	}

	return seeds, maps
}
