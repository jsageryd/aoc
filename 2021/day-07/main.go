package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)

	values := stringsToInts(strings.Split(input, ","))

	fmt.Printf("Part 1: %d\n", alignmentCost(values, median(values)))
}

func median(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sort.Ints(values)
	return values[len(values)/2]
}

func alignmentCost(values []int, alignAt int) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	var cost int

	for _, v := range values {
		cost += abs(v - alignAt)
	}

	return cost
}

func stringsToInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}
	return ints
}
