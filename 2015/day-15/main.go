package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	list := parseInput(input)

	var bestScore int

	distribute(len(list), 100, func(dist []int) {
		var capacity, durability, flavor, texture int

		for idx, amount := range dist {
			capacity += list[idx].capacity * amount
			durability += list[idx].durability * amount
			flavor += list[idx].flavor * amount
			texture += list[idx].texture * amount
		}

		score := max(0, capacity) *
			max(0, durability) *
			max(0, flavor) *
			max(0, texture)

		if score > bestScore {
			bestScore = score
		}
	})

	return bestScore
}

func part2(input []string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	list := parseInput(input)

	var bestScore int

	distribute(len(list), 100, func(dist []int) {
		var capacity, durability, flavor, texture, calories int

		for idx, amount := range dist {
			capacity += list[idx].capacity * amount
			durability += list[idx].durability * amount
			flavor += list[idx].flavor * amount
			texture += list[idx].texture * amount
			calories += list[idx].calories * amount
		}

		if calories != 500 {
			return
		}

		score := max(0, capacity) *
			max(0, durability) *
			max(0, flavor) *
			max(0, texture)

		if score > bestScore {
			bestScore = score
		}
	})

	return bestScore
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(input []string) []ingredient {
	var list []ingredient

	for _, line := range input {
		var i ingredient

		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", "")

		fmt.Sscanf(
			line, "%s capacity %d durability %d flavor %d texture %d calories %d",
			&i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories,
		)

		list = append(list, i)
	}

	return list
}

// distribute generates all possible distributions of n values such that the
// total always adds up to max, calling f with the value of each. For example,
// for n = 3, max = 10, these are some possible distributions:
//
// [10,0,0] (0000000000)
// [9,1,0]  (0000000001)
// [3,3,4]  (0001112222)
func distribute(n, max int, f func(dist []int)) {
	dist := make([]int, n)

	var rec func(idx, max int)

	rec = func(idx, max int) {
		for i := max; i >= 0; i-- {
			dist[idx] = i
			if idx < len(dist)-1 {
				rec(idx+1, max-i)
			}
			if idx == len(dist)-1 {
				f(dist)
				return
			}
		}
	}

	rec(0, max)
}
