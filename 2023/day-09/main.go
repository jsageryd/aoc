package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	for _, line := range input {
		sum += nextValue(line)
	}

	return sum
}

func part2(input []string) int {
	var sum int

	for _, line := range input {
		sum += prevValue(line)
	}

	return sum
}

func nextValue(line string) int {
	seqs := sequences(line)

	for n := len(seqs) - 2; n >= 0; n-- {
		seqs[n] = append(seqs[n], seqs[n][len(seqs[n])-1]+seqs[n+1][len(seqs[n+1])-1])
	}

	return seqs[0][len(seqs[0])-1]
}

func prevValue(line string) int {
	seqs := sequences(line)

	for n := len(seqs) - 2; n >= 0; n-- {
		seqs[n] = append([]int{seqs[n][0] - seqs[n+1][0]}, seqs[n]...)
	}

	return seqs[0][0]
}

func sequences(line string) [][]int {
	var seq []int

	for _, s := range strings.Fields(line) {
		i, _ := strconv.Atoi(s)
		seq = append(seq, i)
	}

	seqs := [][]int{seq}

	for {
		var seq []int
		lastSeq := seqs[len(seqs)-1]
		for i := 1; i < len(lastSeq); i++ {
			seq = append(seq, lastSeq[i]-lastSeq[i-1])
		}
		if len(seq) < 1 {
			break
		}
		seqs = append(seqs, seq)
	}

	return seqs
}
