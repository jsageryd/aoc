package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	count := len(input)*2 + len(input[0])*2 - 4

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if visible(input, x, y) {
				count++
			}
		}
	}

	return count
}

func visible(input []string, x, y int) bool {
	height := func(x, y int) int {
		n, _ := strconv.Atoi(string(input[y][x]))
		return n
	}

	h := height(x, y)

	{
		v := true
		for xx := 0; xx < x; xx++ {
			if height(xx, y) >= h {
				v = false
				break
			}
		}
		if v {
			return true
		}
	}

	{
		v := true
		for xx := x + 1; xx < len(input[y]); xx++ {
			if height(xx, y) >= h {
				v = false
				break
			}
		}
		if v {
			return true
		}
	}

	{
		v := true
		for yy := 0; yy < y; yy++ {
			if height(x, yy) >= h {
				v = false
				break
			}
		}
		if v {
			return true
		}
	}

	{
		v := true
		for yy := y + 1; yy < len(input); yy++ {
			if height(x, yy) >= h {
				v = false
				break
			}
		}
		if v {
			return true
		}
	}

	return false
}
