package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	applyMask := func(val, maskZeros, maskOnes int) int {
		val &= ^maskZeros
		val |= maskOnes
		return val
	}

	var sum int
	mem := make(map[int]int)

	var mZeros, mOnes int

	for _, row := range input {
		switch {
		case strings.HasPrefix(row, "mask"):
			mZeros, mOnes = parseMask(strings.TrimPrefix(row, "mask = "))
		case strings.HasPrefix(row, "mem"):
			var addr, val int
			fmt.Sscanf(row, "mem[%d] = %d", &addr, &val)
			mem[addr] = applyMask(val, mZeros, mOnes)
		}
	}

	for _, v := range mem {
		sum += v
	}

	return sum
}

func parseMask(mask string) (zeros, ones int) {
	for i := range mask {
		switch mask[len(mask)-i-1] {
		case '0':
			zeros += 1 << i
		case '1':
			ones += 1 << i
		}
	}
	return zeros, ones
}
