package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, n)
	}

	fmt.Printf("Part 1: %d\n", findSum(input, 2, 2020))
	fmt.Printf("Part 2: %d\n", findSum(input, 3, 2020))
}

func findSum(list []int, entries int, wantSum int) (prod int) {
	combinations(list, entries, func(comb []int) bool {
		var sum int
		for _, n := range comb {
			sum += n
		}
		if sum == wantSum {
			prod = comb[0]
			for _, n := range comb[1:] {
				prod *= n
			}
			return false
		}
		return true
	})
	return prod
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
