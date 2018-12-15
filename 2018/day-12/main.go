package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	initial := strings.TrimPrefix(scanner.Text(), "initial state: ")

	scanner.Scan()

	var rules []string

	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	state, offset := step(initial, rules, 20)
	fmt.Printf("Part 1: %d\n", sumOfIndicesWithPots(state, offset))

	state, offset = step(initial, rules, 50000000000)
	fmt.Printf("Part 2: %d\n", sumOfIndicesWithPots(state, offset))
}

func sumOfIndicesWithPots(state string, offset int) int {
	sum := 0
	for n := range []byte(state) {
		if state[n] == '#' {
			sum += n - offset
		}
	}
	return sum
}

func step(state string, rules []string, steps int) (string, int) {
	next, offset := recenterAndAdjustLength(stateStringToInternal(state), 0)

	bRules := make([]byte, 0, len(rules))
	for n := range rules {
		if rules[n][9] == '#' {
			bRules = append(bRules, ruleToInternal(rules[n]))
		}
	}

	var lastOffset int
	var prev []byte

	for i := 0; i < steps; i++ {
		if bytes.Equal(prev, next) {
			return internalToStateString(prev), offset + (steps-i)*(offset-lastOffset)
		}
		if len(prev) > len(next) {
			prev = prev[:len(next)]
		} else {
			prev = append(prev, make([]byte, len(next)-len(prev))...)
		}
		copy(prev, next)
	nextFlower:
		for n := range prev {
			next[n] = 0
			for m := range bRules {
				if prev[n] == bRules[m] {
					next[n] = 4
					continue nextFlower
				}
			}
		}
		lastOffset = offset
		next, offset = recenterAndAdjustLength(resultToInternal(next), offset)
	}

	return internalToStateString(next), offset
}

func recenterAndAdjustLength(result []byte, offset int) ([]byte, int) {
	padding := 5
	idx, lidx := -1, -1
	for n := range result {
		if result[n] != 0 {
			idx = n
			break
		}
	}
	for n := len(result) - 1; n >= 0; n-- {
		if result[n] != 0 {
			lidx = n
			break
		}
	}
	var growBy int
	if idx == -1 {
		growBy = padding - len(result)
		offset = (len(result) + growBy) / 2
	} else {
		growBy = padding - idx + (padding - (len(result) - lidx) + 1)
	}
	if growBy > 0 {
		zeros := make([]byte, growBy)
		result = append(result, zeros...)
	}
	if idx != -1 {
		offset += padding - idx
		if idx < padding {
			copy(result[padding-idx:], result)
		} else {
			copy(result[padding:], result[idx:])
		}
		copy(result, make([]byte, padding))
		copy(result[len(result)-padding:], make([]byte, padding))
	}
	if growBy < 0 {
		result = result[:len(result)+growBy]
	}
	return result, offset
}

func resultToInternal(result []byte) []byte {
	for n := 0; n < len(result); n++ {
		var b byte
		for s := uint8(0); s < 5; s++ {
			m := n - 2 + int(s)
			if m >= 0 && m < len(result) && result[m]&4 == 4 {
				b |= 1 << (4 - s)
			}
		}
		result[n] = b
	}
	return result
}

func stateStringToInternal(state string) []byte {
	bState := []byte(state)
	bb := make([]byte, len(bState))
	for n := 0; n < len(bState); n++ {
		var b byte
		for s := uint8(0); s < 5; s++ {
			m := n - 2 + int(s)
			if m >= 0 && m < len(bState) && bState[m] == '#' {
				b |= 1 << (4 - s)
			}
		}
		bb[n] = b
	}
	return bb
}

func internalToStateString(bb []byte) string {
	bState := make([]byte, len(bb))
	for n := range bb {
		if bb[n]&4 == 4 {
			bState[n] = '#'
		} else {
			bState[n] = '.'
		}
	}
	return string(bState)
}

func ruleToInternal(rule string) byte {
	bRule := []byte(rule)
	var b byte
	for s := uint8(0); s < 5; s++ {
		if bRule[s] == '#' {
			b |= 1 << (4 - s)
		}
	}
	return b
}
