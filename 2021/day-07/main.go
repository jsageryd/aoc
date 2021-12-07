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

	fmt.Printf("Part 1: %d\n", part1(values))
	fmt.Printf("Part 2: %d\n", part2(values))
}

func part1(values []int) int {
	return alignmentCost(values, median(values))
}

func part2(values []int) int {
	if len(values) == 0 {
		return 0
	}

	sort.Ints(values)
	min := values[0]
	max := values[len(values)-1]

	minCost := alignmentCost2(values, min)

	for pos := min + 1; pos <= max; pos++ {
		cost := alignmentCost2(values, pos)
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func median(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sort.Ints(values)
	return values[len(values)/2]
}

func alignmentCost(values []int, alignAt int) int {
	var cost int

	for _, v := range values {
		cost += abs(v - alignAt)
	}

	return cost
}

func alignmentCost2(values []int, alignAt int) int {
	var cost int

	for _, v := range values {
		for i := 1; i <= abs(v-alignAt); i++ {
			cost += i
		}
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
