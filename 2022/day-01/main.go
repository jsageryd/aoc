package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", maxCalories(input))
}

func maxCalories(input []string) int {
	var cur, max int

	for _, line := range input {
		if line == "" {
			cur = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		cur += cal

		if cur > max {
			max = cur
		}
	}

	return max
}
