package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", highestSeatID(input))
	fmt.Printf("Part 2: %d\n", findMySeatID(input))
}

func highestSeatID(specs []string) int {
	var highestID int

	for _, spec := range specs {
		_, _, id := parseSpec(spec)
		if id > highestID {
			highestID = id
		}
	}

	return highestID
}

func findMySeatID(specs []string) int {
	ids := make([]int, 0, len(specs))
	for _, spec := range specs {
		_, _, id := parseSpec(spec)
		ids = append(ids, id)
	}
	return findMissing(ids)
}

func findMissing(s []int) int {
	sort.Ints(s)
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] > 1 {
			return s[i] - 1
		}
	}
	return 0
}

func parseSpec(spec string) (row, col, id int) {
	parseBase2 := func(spec string, one rune) int {
		var res int
		p := 1 << (len(spec) - 1)
		for _, r := range spec {
			if r == one {
				res += p
			}
			p /= 2
		}
		return res
	}

	row = parseBase2(spec[:7], 'B')
	col = parseBase2(spec[7:], 'R')

	return row, col, row*8 + col
}
