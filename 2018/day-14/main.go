package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr string

	fmt.Scan(&inputStr)

	input, _ := strconv.Atoi(inputStr)

	var scores []string
	for _, s := range scoresAfterNth(input, 10) {
		scores = append(scores, strconv.Itoa(s))
	}

	fmt.Printf("Part 1: %s\n", strings.Join(scores, ""))

	seq := make([]int, len(inputStr))
	for n := range inputStr {
		seq[n], _ = strconv.Atoi(string(inputStr[n]))
	}

	fmt.Printf("Part 2: %d\n", numberOfRecipesBeforeSequence(seq))
}

func scoresAfterNth(n int, count int) []int {
	recipes := make([]int, 0, n+count)
	c := makeRecipes()
	for i := 0; i < n+count; i++ {
		recipes = append(recipes, <-c)
	}
	return recipes[n:]
}

func numberOfRecipesBeforeSequence(seq []int) int {
	c := makeRecipes()
	buf := make([]int, len(seq))
	var count int
	for n := range seq {
		buf[n] = <-c
	}
	var found bool
	for !found {
		found = true
		for n := range buf {
			if buf[n] != seq[n] {
				found = false
			}
		}
		copy(buf, buf[1:])
		buf[len(buf)-1] = <-c
		count++
	}
	return count - 1
}

func makeRecipes() <-chan int {
	c := make(chan int, 1000)

	go func() {
		elf1, elf2 := 0, 1
		recipes := []int{3, 7}
		c <- 3
		c <- 7
		for {
			sum := recipes[elf1] + recipes[elf2]
			if sum > 9 {
				recipes = append(recipes, sum/10)
				c <- recipes[len(recipes)-1]
			}
			recipes = append(recipes, sum%10)
			c <- recipes[len(recipes)-1]
			elf1 = (elf1 + (1 + recipes[elf1])) % len(recipes)
			elf2 = (elf2 + (1 + recipes[elf2])) % len(recipes)
		}
	}()

	return c
}
