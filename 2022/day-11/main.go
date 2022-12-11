package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	monkeys := parseMonkeys(input)

	for n := 0; n < 20; n++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inspectAndThrow(func(worry *int) {
					*worry /= 3
				})
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})

	return monkeys[0].itemsInspected * monkeys[1].itemsInspected
}

func part2(input []string) int {
	monkeys := parseMonkeys(input)

	div := 1
	for _, m := range monkeys {
		div *= m.testDiv
	}

	for n := 0; n < 10000; n++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inspectAndThrow(func(worry *int) {
					*worry %= div
				})
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})

	return monkeys[0].itemsInspected * monkeys[1].itemsInspected
}

func parseMonkeys(input []string) []*monkey {
	var monkeys []*monkey

	for n := range input {
		if !strings.HasPrefix(input[n], "Monkey") {
			continue
		}

		var m monkey

		// Starting items
		itemsStr := strings.TrimPrefix(input[n+1], "  Starting items: ")
		for _, s := range regexp.MustCompile(", ?").Split(itemsStr, -1) {
			item, _ := strconv.Atoi(s)
			m.items = append(m.items, item)
		}

		// Operation
		m.op = strings.TrimPrefix(input[n+2], "  Operation: new = ")

		// Test
		fmt.Sscanf(input[n+3], " Test: divisible by %d", &m.testDiv)

		// If true
		fmt.Sscanf(input[n+4], " If true: throw to monkey %d", &m.trueMonkey)

		// If false
		fmt.Sscanf(input[n+5], " If false: throw to monkey %d", &m.falseMonkey)

		monkeys = append(monkeys, &m)
	}

	for _, m := range monkeys {
		m.friends = monkeys
	}

	return monkeys
}

type monkey struct {
	items          []int
	op             string
	testDiv        int
	trueMonkey     int
	falseMonkey    int
	friends        []*monkey
	itemsInspected int
}

func (m *monkey) inspectAndThrow(manageWorry func(worry *int)) {
	if len(m.items) == 0 {
		return
	}

	item := m.items[0]
	m.items = m.items[1:]

	var left, right int
	op := strings.Fields(m.op)
	if op[0] == "old" {
		left = item
	} else {
		left, _ = strconv.Atoi(op[0])
	}
	if op[2] == "old" {
		right = item
	} else {
		right, _ = strconv.Atoi(op[2])
	}
	switch op[1] {
	case "*":
		item = left * right
	case "+":
		item = left + right
	}

	m.itemsInspected++

	manageWorry(&item)

	var nextMonkey *monkey

	if item%m.testDiv == 0 {
		nextMonkey = m.friends[m.trueMonkey]
	} else {
		nextMonkey = m.friends[m.falseMonkey]
	}

	nextMonkey.items = append(nextMonkey.items, item)
}
