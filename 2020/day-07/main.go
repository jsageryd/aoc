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
}

func part1(rules []string) int {
	m := parseRules(rules)

	findParents := func(child string) []string {
		var parents []string
		for parentBag, childBags := range m {
			for _, childBag := range childBags {
				if childBag == child {
					parents = append(parents, parentBag)
					break
				}
			}
		}
		return parents
	}

	bags := make(map[string]struct{})

	stack := findParents("shiny gold")
	for len(stack) > 0 {
		bag := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := bags[bag]; !ok {
			bags[bag] = struct{}{}
			stack = append(stack, findParents(bag)...)
		}
	}

	return len(bags)
}

// parseRules parses rules into a map of bag -> list of bags
func parseRules(rules []string) map[string][]string {
	m := make(map[string][]string)
	for _, rule := range rules {
		parentBag, childBags := parseRule(rule)
		if len(childBags) > 0 {
			m[parentBag] = childBags
		}
	}
	return m
}

func parseRule(rule string) (parentBag string, childBags []string) {
	split := strings.SplitN(rule, " ", 5)
	parentBag = strings.Join(split[:2], " ")

	for _, s := range strings.Split(split[len(split)-1], ", ") {
		var count int
		var c1, c2 string
		fmt.Sscanf(s, "%d %s %s", &count, &c1, &c2)
		for i := 0; i < count; i++ {
			childBags = append(childBags, c1+" "+c2)
		}
	}

	return parentBag, childBags
}
