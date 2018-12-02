package main

import (
	"fmt"
)

func main() {
	var input []string

	for {
		var id string
		if _, err := fmt.Scanln(&id); err != nil {
			break
		}
		input = append(input, id)
	}

	fmt.Printf("Part 1: %d\n", checksum(input))
	fmt.Printf("Part 2: %s\n", commonLettersOfCorrectBoxes(input))
}

func checksum(ids []string) int {
	twos, threes := 0, 0

	for _, id := range ids {
		freq := map[rune]int{}
		for _, r := range id {
			freq[r]++
		}
		var two, three bool
		for _, f := range freq {
			switch f {
			case 2:
				two = true
			case 3:
				three = true
			}
		}
		if two {
			twos++
		}
		if three {
			threes++
		}
	}

	return twos * threes
}

func commonLettersOfCorrectBoxes(ids []string) string {
	for i, id1 := range ids {
		for j, id2 := range ids {
			if i == j {
				continue
			}
			var common string
			for n := range id1 {
				if id1[n] == id2[n] {
					common = common + string(id1[n])
				}
			}
			if len(common) == len(id1)-1 {
				return common
			}
		}
	}
	return ""
}
