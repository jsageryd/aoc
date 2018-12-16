package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int

	fmt.Scan(&input)

	var scores []string
	for _, s := range scoresAfterNth(input, 10) {
		scores = append(scores, strconv.Itoa(s))
	}

	fmt.Printf("Part 1: %s\n", strings.Join(scores, ""))
}

func scoresAfterNth(n int, count int) []int {
	return makeRecipes(n + count)[n:]
}

func makeRecipes(n int) []int {
	elf1 := 0
	elf2 := 1
	recipes := []int{3, 7}
	for len(recipes) < n {
		sum := recipes[elf1] + recipes[elf2]
		if sum > 9 {
			recipes = append(recipes, sum/10)
		}
		recipes = append(recipes, sum%10)
		elf1 = (elf1 + (1 + recipes[elf1])) % len(recipes)
		elf2 = (elf2 + (1 + recipes[elf2])) % len(recipes)
	}
	return recipes[:n]
}
