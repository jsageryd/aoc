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
}

func part1(input []string) int {
	var habs []handAndBid

	for _, line := range input {
		var hab handAndBid
		fmt.Sscanf(line, "%s %d", &hab.hand, &hab.bid)
		habs = append(habs, hab)
	}

	slices.SortFunc(habs, func(a, b handAndBid) int {
		if v := cmp.Compare(handType(a.hand), handType(b.hand)); v != 0 {
			return v
		}

		for n := range a.hand {
			if v := cmp.Compare(cardRank(a.hand[n]), cardRank(b.hand[n])); v != 0 {
				return v
			}
		}

		return 0
	})

	slices.Reverse(habs)

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

func handType(hand string) int {
	freq := make(map[byte]int)

	for n := range hand {
		freq[hand[n]]++
	}

	var id int

	var freqs []int

	for _, f := range freq {
		freqs = append(freqs, f)
	}

	slices.Sort(freqs)

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

func cardRank(card byte) int {
	return strings.IndexByte("AKQJT98765432", card)
}
