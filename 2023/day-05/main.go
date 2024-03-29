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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []byte) int {
	seeds, maps := parse(input)

	var locations []int

	for _, cur := range seeds {

	nextMap:
		for _, m := range maps {
			for _, l := range m {
				if cur >= l[1] && cur < l[1]+l[2] {
					cur = l[0] + cur - l[1]
					continue nextMap
				}
			}
		}

		locations = append(locations, cur)
	}

	return slices.Min(locations)
}

func part2(input []byte) int {
	seeds, maps := parse(input)

	var minLocation int

	for n := 0; n < len(seeds)-1; n += 2 {
		start, length := seeds[n], seeds[n+1]

		for seed := start; seed < start+length; seed++ {
			cur := seed

		nextMap:
			for _, m := range maps {
				for _, l := range m {
					if cur >= l[1] && cur < l[1]+l[2] {
						cur = l[0] + cur - l[1]
						continue nextMap
					}
				}
			}

			if minLocation == 0 || cur < minLocation {
				minLocation = cur
			}
		}
	}

	return minLocation
}

func parse(input []byte) (seeds []int, maps [][][3]int) {
	parts := strings.Split(string(input), "\n\n")

	for _, s := range strings.Fields(parts[0][7:]) {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}

	for _, mapStr := range parts[1:] {
		var lines [][3]int

		for _, lineStr := range strings.Split(strings.TrimSpace(mapStr), "\n")[1:] {
			var l [3]int
			fmt.Sscanf(lineStr, "%d %d %d", &l[0], &l[1], &l[2])
			lines = append(lines, l)
		}

		maps = append(maps, lines)
	}

	return seeds, maps
}
