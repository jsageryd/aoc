package main

import (
	"bufio"
	"cmp"
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
	habs := parse(input)

	sortHandAndBids(habs, false)

	return totalWinnings(habs)
}

func part2(input []string) int {
	habs := parse(input)

	sortHandAndBids(habs, true)

	return totalWinnings(habs)
}

func parse(input []string) []handAndBid {
	var habs []handAndBid

	for _, line := range input {
		var hab handAndBid
		fmt.Sscanf(line, "%s %d", &hab.hand, &hab.bid)
		habs = append(habs, hab)
	}

	return habs
}

func sortHandAndBids(habs []handAndBid, jokerRule bool) {
	slices.SortFunc(habs, func(a, b handAndBid) int {
		if v := cmp.Compare(
			handType(a.hand, jokerRule),
			handType(b.hand, jokerRule),
		); v != 0 {
			return v
		}

		for n := range a.hand {
			if v := cmp.Compare(
				cardRank(a.hand[n], jokerRule),
				cardRank(b.hand[n], jokerRule),
			); v != 0 {
				return v
			}
		}

		return 0
	})

	slices.Reverse(habs)
}

func totalWinnings(habs []handAndBid) int {
	var sum int

	for n := range habs {
		sum += habs[n].bid * (n + 1)
	}

	return sum
}

type handAndBid struct {
	hand string
	bid  int
}

func handType(hand string, jokerRule bool) int {
	freq := make(map[byte]int)

	for n := range hand {
		freq[hand[n]]++
	}

	var id int

	var freqs []int

	if jokerRule && freq['J'] == 5 {
		return 0
	}

	for c, f := range freq {
		if jokerRule && c == 'J' {
			continue
		}
		freqs = append(freqs, f)
	}

	slices.Sort(freqs)

	if jokerRule {
		freqs[len(freqs)-1] += freq['J']
	}

	m := 1
	for n := len(freqs) - 1; n >= 0; n-- {
		id += freqs[n] * m
		m *= 10
	}

	return map[int]int{
		5:     0, // Five of a kind
		14:    1, // Four of a kind
		23:    2, // Full house
		113:   3, // Three of a kind
		122:   4, // Two pair
		1112:  5, // One pair
		11111: 6, // High card
	}[id]
}

func cardRank(card byte, jokerRule bool) int {
	if jokerRule {
		return strings.IndexByte("AKQT98765432J", card)
	} else {
		return strings.IndexByte("AKQJT98765432", card)
	}
}
