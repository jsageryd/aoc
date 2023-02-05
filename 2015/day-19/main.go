package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	rs, s := parse(input)

	m := make(map[string]struct{})

	for _, r := range rs {
		for _, idx := range indices(s, r.from) {
			m[s[:idx]+r.to+s[idx+len(r.from):]] = struct{}{}
		}
	}

	return len(m)
}

func part2(input []string) int {
	const limit = 100

	for tries := 0; tries < limit; tries++ {
		rs, s := parse(input)

		rand.Shuffle(len(rs), func(i, j int) {
			rs[i], rs[j] = rs[j], rs[i]
		})

		var count int
		var ok bool

		for {
			if s, ok = reduce(rs, s); !ok {
				break
			}
			count++
		}

		if s == "e" {
			return count
		}
	}

	return 0
}

type replacement struct {
	from, to string
}

func parse(input []string) (rs []replacement, s string) {
	for _, line := range input[:len(input)-2] {
		from, to, _ := strings.Cut(line, " => ")
		rs = append(rs, replacement{from: from, to: to})
	}

	return rs, input[len(input)-1]
}

func indices(s, sep string) []int {
	var idxs []int

	for n := 0; n <= len(s); n++ {
		if strings.HasPrefix(s[n:], sep) {
			idxs = append(idxs, n)
		}
	}

	return idxs
}

func reduce(rs []replacement, s string) (string, bool) {
	for _, r := range rs {
		if strings.Contains(s, r.to) {
			return strings.Replace(s, r.to, r.from, 1), true
		}
	}
	return s, false
}
