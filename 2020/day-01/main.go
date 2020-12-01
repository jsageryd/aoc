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

	n1, n2 := findSum(input, 2020)

	fmt.Printf("Part 1: %d\n", n1*n2)
}

func findSum(list []int, wantSum int) (n1, n2 int) {
	for _, n1 := range list {
		for _, n2 := range list {
			if n1+n2 == wantSum {
				return n1, n2
			}
		}
	}
	return 0, 0
}
