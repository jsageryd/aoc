package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var count int
	for _, entry := range input {
		sep := strings.IndexByte(entry, '|')
		values := strings.Fields(entry[sep:])
		for _, v := range values {
			switch len(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func part2(input []string) int {
	var sum int

	for _, entry := range input {
		sep := strings.IndexByte(entry, '|')
		patterns := strings.Fields(entry[:sep])
		values := strings.Fields(entry[sep+1:])

		if len(values) != 4 {
			log.Fatalf("unexpected number of values: %d", len(values))
		}

		sort.Slice(patterns, func(i, j int) bool {
			return len(patterns[i]) < len(patterns[j])
		})

		sortRunesInStrings(patterns)

		p2d := make(map[string]int)
		d2p := make(map[int]string)

		for _, p := range patterns {
			if d, ok := patternToDigit(p, d2p); ok {
				p2d[p], d2p[d] = d, p
			}
		}

		if len(d2p) != 10 {
			log.Fatalf("only managed to map %d digits", len(d2p))
		}

		sortRunesInStrings(values)

		sum += p2d[values[0]]*1000 +
			p2d[values[1]]*100 +
			p2d[values[2]]*10 +
			p2d[values[3]]
	}

	return sum
}

func patternToDigit(pattern string, d2p map[int]string) (digit int, ok bool) {
	switch len(pattern) {
	case 2: // digit is 1
		return 1, true
	case 3: // digit is 7
		return 7, true
	case 4: // digit is 4
		return 4, true
	case 5: // digit is 2, 3, or 5
		if one, ok := d2p[1]; ok && containsAll(pattern, []rune(one)) {
			return 3, true
		}
		if four, ok := d2p[4]; ok {
			switch containsCount(pattern, []rune(four)) {
			case 2:
				return 2, true
			case 3:
				return 5, true
			}
		}
	case 6: // digit is 0, 6, or 9
		if four, ok := d2p[4]; ok && containsAll(pattern, []rune(four)) {
			return 9, true
		}
		if one, ok := d2p[1]; ok && containsAll(pattern, []rune(one)) {
			return 0, true
		}
		return 6, true
	case 7: // digit is 8
		return 8, true
	}

	return 0, false
}

func sortRunesInStrings(strs []string) {
	for i := range strs {
		runes := []rune(strs[i])
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		strs[i] = string(runes)
	}
}

func containsAll(str string, runes []rune) bool {
	for _, r := range runes {
		if !strings.ContainsRune(str, r) {
			return false
		}
	}
	return true
}

func containsCount(str string, runes []rune) int {
	var count int
	for _, r := range runes {
		if strings.ContainsRune(str, r) {
			count++
		}
	}
	return count
}
