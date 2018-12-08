package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	score, garbage := scoreAndGarbage(input)

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", garbage)
}

func scoreAndGarbage(input []byte) (score int, garbage int) {
	level := 0
	inGarbage := false

	for b := 0; b < len(input); b++ {
		if inGarbage {
			switch input[b] {
			case '!':
				b++
			case '>':
				inGarbage = false
			default:
				garbage++
			}
		} else {
			switch input[b] {
			case '{':
				level++
				score += level
			case '}':
				level--
			case '<':
				inGarbage = true
			}
		}
	}

	return score, garbage
}
