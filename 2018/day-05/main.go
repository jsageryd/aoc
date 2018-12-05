package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)

	fmt.Printf("Part 1: %d\n", len(reduce(input)))
	fmt.Printf("Part 2: %d\n", len(reduceAlternate(input)))
}

func reduce(polymer string) string {
	units := []byte(polymer)

	for loop := true; loop; {
		loop = false
		for n := 0; n < len(units)-1; n++ {
			if units[n] == '.' {
				continue
			}
			m := n + 1
			for ; m < len(units) && units[m] == '.'; m++ {
			}
			if m >= len(units) {
				break
			}
			if units[n] != units[m] && asciiToLower(units[n]) == asciiToLower(units[m]) {
				units[n], units[m] = '.', '.'
				loop = true
			}
		}
	}

	return strings.Replace(string(units), ".", "", -1)
}

func reduceAlternate(polymer string) string {
	m := map[byte]struct{}{}

	for _, b := range []byte(polymer) {
		m[asciiToLower(b)] = struct{}{}
	}

	shortest := polymer

	for unit := range m {
		p := []byte(polymer)
		for n := range p {
			if asciiToLower(p[n]) == unit {
				p[n] = '.'
			}
		}
		if pp := reduce(string(p)); len(pp) < len(shortest) {
			shortest = pp
		}
	}

	return shortest
}

func asciiToLower(a byte) byte {
	if a&0x20 == 0 {
		return a ^ 0x20
	}
	return a
}
