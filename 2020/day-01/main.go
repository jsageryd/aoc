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

	n1, n2 := findSumOfTwo(input, 2020)
	fmt.Printf("Part 1: %d\n", n1*n2)

	n1, n2, n3 := findSumOfThree(input, 2020)
	fmt.Printf("Part 2: %d\n", n1*n2*n3)
}

func findSumOfTwo(list []int, wantSum int) (n1, n2 int) {
	for _, n1 := range list {
		for _, n2 := range list {
			if n1+n2 == wantSum {
				return n1, n2
			}
		}
	}
	return 0, 0
}

func findSumOfThree(list []int, wantSum int) (n1, n2, n3 int) {
	for _, n1 := range list {
		for _, n2 := range list {
			for _, n3 := range list {
				if n1+n2+n3 == wantSum {
					return n1, n2, n3
				}
			}
		}
	}
	return 0, 0, 0
}
