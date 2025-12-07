package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var input []int

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		v, _ := strconv.Atoi(s.Text())
		input = append(input, v)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []int) int {
	var sum int

	for _, v := range input {
		sum += v
	}

	targetWeight := sum / 3

	slices.SortFunc(input, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	var groups [][]int

	combTargetSum(input, targetWeight, func(comb1 []int) bool {
		combTargetSum(subtract(input, comb1), targetWeight, func(comb2 []int) bool {
			comb3 := subtract(subtract(input, comb1), comb2)
			groups = [][]int{comb1, comb2, comb3}
			return false
		})
		return false
	})

	slices.SortFunc(groups, func(a, b []int) int {
		if lCmp := cmp.Compare(len(a), len(b)); lCmp != 0 {
			return lCmp
		}
		return cmp.Compare(quantumEntanglement(a), quantumEntanglement(b))
	})

	return quantumEntanglement(groups[0])
}

func quantumEntanglement(group []int) int {
	m := 1
	for _, v := range group {
		m *= v
	}
	return m
}

func subtract(a, b []int) []int {
	var result []int

	for _, v1 := range a {
		if !slices.Contains(b, v1) {
			result = append(result, v1)
		}
	}

	return result
}

// combTargetSum picks elements from s until they sum to the given target value
// and calls f for each combination of elements. If f returns false, the search
// stops.
func combTargetSum(s []int, target int, f func(comb []int) bool) {
	cont := true

	var rec func(ss []int, cc []int, target int)

	rec = func(ss []int, cc []int, target int) {
		for n := range ss {
			if !cont {
				break
			}

			if ss[n] < target {
				rec(ss[n+1:], append(cc, ss[n]), target-ss[n])
			} else if ss[n] == target {
				cont = cont && f(append(cc, ss[n]))
			}
		}
	}

	rec(s, nil, target)
}
