package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", checksum(input))
	fmt.Printf("Part 2: %d\n", checksumDivisibles(input))
}

func checksum(input []byte) int {
	sum := 0

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		var nums []int
		for _, s := range strings.Split(scanner.Text(), "\t") {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		lineMin, lineMax := nums[0], nums[0]
		for _, n := range nums {
			switch {
			case n < lineMin:
				lineMin = n
			case n > lineMax:
				lineMax = n
			}
		}
		sum += lineMax - lineMin
	}

	return sum
}

func checksumDivisibles(input []byte) int {
	sum := 0

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		var nums []int
		for _, s := range strings.Split(scanner.Text(), "\t") {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
	outer:
		for i, n := range nums {
			for j, m := range nums {
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
