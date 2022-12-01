package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	first, second, third := maxCalories(input)

	fmt.Printf("Part 1: %d\n", first)
	fmt.Printf("Part 2: %d\n", first+second+third)
}

func maxCalories(input []string) (first, second, third int) {
	var cur int
	var sums []int

	for _, line := range input {
		if line == "" {
			sums = append(sums, cur)
			cur = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		cur += cal
	}

	sums = append(sums, cur)

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums[0], sums[1], sums[2]
}
