package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	base2 := func(r rune) rune {
		switch r {
		case 'F', 'L':
			return '0'
		case 'B', 'R':
			return '1'
		default:
			return r
		}
	}

	r, _ := strconv.ParseInt(strings.Map(base2, spec[:7]), 2, 64)
	c, _ := strconv.ParseInt(strings.Map(base2, spec[7:]), 2, 64)

	row, col = int(r), int(c)

	return row, col, row*8 + col
}
