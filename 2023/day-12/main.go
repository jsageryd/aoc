package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

	for _, row := range input {
		sum += arrangements(row)
	}

	return sum
}

func arrangements(row string) int {
	var count int

	var groups []int

	var totalBroken int

	pattern, groupsStr, _ := strings.Cut(row, " ")
	for _, s := range strings.Split(groupsStr, ",") {
		group, _ := strconv.Atoi(s)
		groups = append(groups, group)
		totalBroken += group
	}

	b := bytes.Repeat([]byte{'#'}, totalBroken)
	b = append(b, bytes.Repeat([]byte{'.'}, len(pattern)-totalBroken)...)

	if valid(pattern, groups, b) {
		count++
	}

	for perm(b) {
		if valid(pattern, groups, b) {
			count++
		}
	}

	return count
}

func valid(pattern string, groups []int, b []byte) bool {
	if len(pattern) != len(b) {
		return false
	}

	for n := range pattern {
		if pattern[n] != '?' && pattern[n] != b[n] {
			return false
		}
	}

	var gotGroups []int
	var curGroup int

	for n := range b {
		switch b[n] {
		case '#':
			curGroup++
		case '.':
			if curGroup > 0 {
				gotGroups = append(gotGroups, curGroup)
				curGroup = 0
			}
		}
	}

	if curGroup > 0 {
		gotGroups = append(gotGroups, curGroup)
	}

	if !slices.Equal(groups, gotGroups) {
		return false
	}

	return true
}

// perm permutates b in-place into its next lexicographical permutation and
// returns true if successful, or false if there are no more permutations.
//
// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func perm(b []byte) bool {
	var k int
	for i := len(b) - 2; i >= 0; i-- {
		if b[i] < b[i+1] {
			k = i
			break
		}
	}
	for l := len(b) - 1; l > k; l-- {
		if b[k] < b[l] {
			b[k], b[l] = b[l], b[k]
			rev(b[k+1:])
			return true
		}
	}
	return false
}

// rev reverses the elements of b.
func rev(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
