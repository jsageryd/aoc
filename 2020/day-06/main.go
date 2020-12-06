package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var groups [][]string

	var group []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			groups = append(groups, group)
			group = nil
			continue
		}
		group = append(group, row)
	}
	groups = append(groups, group)

	fmt.Printf("Part 1: %d\n", part1(groups))
	fmt.Printf("Part 2: %d\n", part2(groups))
}

func part1(groups [][]string) int {
	var sum int
	for _, group := range groups {
		m := make(map[rune]struct{})
		for _, answers := range group {
			for _, answer := range answers {
				m[answer] = struct{}{}
			}
		}
		sum += len(m)
	}
	return sum
}

func part2(groups [][]string) int {
	var sum int
	for _, group := range groups {
		m := make(map[rune]struct{})
		for _, answer := range group[0] {
			m[answer] = struct{}{}
		}
		for _, answers := range group {
			m2 := make(map[rune]struct{})
			for _, answer := range answers {
				m2[answer] = struct{}{}
			}
			for haveAnswer := range m {
				if _, ok := m2[haveAnswer]; !ok {
					delete(m, haveAnswer)
				}
			}
		}
		sum += len(m)
	}
	return sum
}
