package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var containers []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, c)
	}

	fmt.Printf("Part 1: %d\n", part1(containers, 150))
}

func part1(containers []int, litres int) int {
	var count int

	for k := 1; k < len(containers); k++ {
		combinations(containers, k, func(comb []int) bool {
			var sum int
			for _, c := range comb {
				sum += c
			}
			if sum == litres {
				count++
			}
			return true
		})
	}

	return count
}

// combinations picks k elements from s and calls f for each combination, until
// there are no more combinations or the given function returns false.
func combinations(s []int, k int, f func(comb []int) bool) {
	cont := true

	comb := make([]int, k)

	var rec func(ss []int, cc []int)

	rec = func(ss []int, cc []int) {
		for n := 0; n <= len(ss)-len(cc) && cont; n++ {
			cc[0] = ss[n]
			if len(cc) > 1 {
				rec(ss[n+1:], cc[1:])
			} else {
				cont = cont && f(comb)
			}
		}
	}

	rec(s, comb)

	/*
		Algorithm example

		abcdef pick 3
		^^^
		^^ ^
		^^  ^
		^^   ^
		^ ^^
		^ ^ ^
		^ ^  ^
		^  ^^
		^  ^ ^
		^   ^^
		 ^^^
		 ^^ ^
		 ^^  ^
		 ^ ^^
		 ^ ^ ^
		 ^  ^^
		  ^^^
		  ^^ ^
		  ^ ^^
		   ^^^
	*/
}
