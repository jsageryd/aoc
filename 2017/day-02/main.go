package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var row []int
		for _, s := range strings.Split(scanner.Text(), "\t") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		rows = append(rows, row)
	}

	fmt.Printf("Part 1: %d\n", checksum(rows))
	fmt.Printf("Part 2: %d\n", checksumDivisibles(rows))
}

func checksum(rows [][]int) int {
	sum := 0

	for _, row := range rows {
		rowMin, rowMax := row[0], row[0]
		for _, n := range row {
			switch {
			case n < rowMin:
				rowMin = n
			case n > rowMax:
				rowMax = n
			}
		}
		sum += rowMax - rowMin
	}

	return sum
}

func checksumDivisibles(rows [][]int) int {
	sum := 0

	for _, row := range rows {
	outer:
		for i, n := range row {
			for j, m := range row {
				if i == j {
					continue
				}
				if n%m == 0 {
					sum += n / m
					break outer
				}
			}
		}
	}

	return sum
}
