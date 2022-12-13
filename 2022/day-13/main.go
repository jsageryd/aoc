package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []byte) int {
	var sum int

	dec := json.NewDecoder(bytes.NewReader(input))

	for idx := 1; dec.More(); idx++ {
		var a, b any

		if err := dec.Decode(&a); err != nil {
			log.Fatal(err)
		}

		if err := dec.Decode(&b); err != nil {
			log.Fatal(err)
		}

		if compare(a, b) <= 0 {
			sum += idx
		}
	}

	return sum
}

func part2(input []byte) int {
	const divider1 = "[[2]]"
	const divider2 = "[[6]]"

	input = append(input, []byte(divider1+"\n")...)
	input = append(input, []byte(divider2+"\n")...)

	dec := json.NewDecoder(bytes.NewReader(input))

	var packets []any

	for dec.More() {
		var p any

		if err := dec.Decode(&p); err != nil {
			log.Fatal(err)
		}

		packets = append(packets, p)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == -1
	})

	var idx1, idx2 int

	for n, p := range packets {
		if fmt.Sprint(p) == divider1 {
			idx1 = n + 1
		}
		if fmt.Sprint(p) == divider2 {
			idx2 = n + 1
		}
	}

	return idx1 * idx2
}

// compare returns -1 if a is less than b, 0 if a is equal to b, and 1 if a is
// greater than b, according to the puzzle rules.
func compare(a, b any) int {
	sliceA, aIsSlice := a.([]any)
	sliceB, bIsSlice := b.([]any)

	if aIsSlice && !bIsSlice {
		sliceB, bIsSlice = []any{b}, true
	}

	if bIsSlice && !aIsSlice {
		sliceA, aIsSlice = []any{a}, true
	}

	if aIsSlice && bIsSlice {
		for n := 0; n < len(sliceA) && n < len(sliceB); n++ {
			if c := compare(sliceA[n], sliceB[n]); c != 0 {
				return c
			}
		}
		if len(sliceA) < len(sliceB) {
			return -1
		}
		if len(sliceA) > len(sliceB) {
			return 1
		}
	} else {
		if a.(float64) < b.(float64) {
			return -1
		}
		if a.(float64) > b.(float64) {
			return 1
		}
	}

	return 0
}
