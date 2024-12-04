package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var sum int

	for d := range 8 {
		for _, line := range rotate(input, d) {
			sum += strings.Count(line, "XMAS")
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int

	xmas := []string{
		"M.M",
		".A.",
		"S.S",
	}

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[0])-1; x++ {
			if input[y][x] == 'A' {
				var block []string

				block = append(block, string(input[y-1][x-1])+"."+string(input[y-1][x+1]))
				block = append(block, ".A.")
				block = append(block, string(input[y+1][x-1])+"."+string(input[y+1][x+1]))

				for _, d := range []int{0, 2, 4, 6} {
					if slices.Equal(xmas, rotate(block, d)) {
						sum++
						break
					}
				}
			}
		}
	}

	return sum
}

// rotate rotates the input counter-clockwise. The lines of the input must have
// equal length. The input is assumed to be ascii.
//
// The direction is a number from 0 to 7:
//
// 0: 0°
// 1: 45°
// 2: 90°
// 3: 135°
// 4: 180°
// 5: 225°
// 6: 270°
// 7: 315°
func rotate(input []string, direction int) []string {
	if len(input) == 0 || len(input[0]) == 0 {
		return input
	}

	var out []string

	switch direction {
	case 0:
		out = input
	case 1:
		for startX := len(input[0]) - 1; startX >= 0; startX-- {
			var line []byte
			for y, x := 0, startX; y < len(input) && x < len(input[0]); y, x = y+1, x+1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
		for startY := 1; startY < len(input); startY++ {
			var line []byte
			for y, x := startY, 0; y < len(input) && x < len(input[0]); y, x = y+1, x+1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 2:
		for x := len(input[0]) - 1; x >= 0; x-- {
			var line []byte
			for y := 0; y < len(input); y++ {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 3:
		for startY := len(input) - 1; startY >= 0; startY-- {
			var line []byte
			for y, x := startY, len(input[0])-1; y < len(input) && x >= 0; y, x = y+1, x-1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
		for startX := len(input[0]) - 2; startX >= 0; startX-- {
			var line []byte
			for y, x := 0, startX; y < len(input) && x >= 0; y, x = y+1, x-1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 4:
		for y := len(input) - 1; y >= 0; y-- {
			var line []byte
			for x := len(input[0]) - 1; x >= 0; x-- {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 5:
		for startX := 0; startX < len(input[0]); startX++ {
			var line []byte
			for y, x := len(input)-1, startX; y >= 0 && x >= 0; y, x = y-1, x-1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
		for startY := len(input) - 2; startY >= 0; startY-- {
			var line []byte
			for y, x := startY, len(input[0])-1; y >= 0 && x >= 0; y, x = y-1, x-1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 6:
		for x := 0; x < len(input[0]); x++ {
			var line []byte
			for y := len(input) - 1; y >= 0; y-- {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	case 7:
		for startY := 0; startY < len(input); startY++ {
			var line []byte
			for y, x := startY, 0; y >= 0 && x < len(input); y, x = y-1, x+1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
		for startX := 1; startX < len(input[0]); startX++ {
			var line []byte
			for y, x := len(input)-1, startX; y >= 0 && x < len(input[0]); y, x = y-1, x+1 {
				line = append(line, input[y][x])
			}
			out = append(out, string(line))
		}
	}

	return out
}
