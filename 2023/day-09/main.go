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
}

func part1(input []string) int {
	var sum int

	for _, line := range input {
		sum += nextValue(line)
	}

	return sum
}

func nextValue(line string) int {
	seqs := [][]int{ints(line)}

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

	for n := len(seqs) - 2; n >= 0; n-- {
		seqs[n] = append(seqs[n], seqs[n][len(seqs[n])-1]+seqs[n+1][len(seqs[n+1])-1])
	}

	return seqs[0][len(seqs[0])-1]
}

func ints(s string) []int {
	var is []int
	for _, ss := range strings.Fields(s) {
		i, _ := strconv.Atoi(ss)
		is = append(is, i)
	}
	return is
}
