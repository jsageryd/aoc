package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	fmt.Printf("Part 1: %s\n", nextPassword(input))
}

func nextPassword(current string) string {
	newPassword := []rune(current)

	for {
		for i := len(newPassword) - 1; i > 0; i-- {
			if newPassword[i] == 'z' {
				newPassword[i] = 'a'
			} else {
				newPassword[i]++
				break
			}
		}
		if valid(string(newPassword)) {
			break
		}
	}

	return string(newPassword)
}

func valid(password string) bool {
	// Rule 1: Passwords must include an increasing straight of three letters.
	rule1 := func(password string) bool {
		var lastR rune
		var count int
		for _, r := range password {
			if r-lastR == 1 {
				count++
			} else {
				count = 1
			}
			if count == 3 {
				return true
			}
			lastR = r
		}
		return false
	}

	// Rule 2: Passwords may not contain i, o, or l.
	rule2 := func(password string) bool {
		return !strings.ContainsAny(password, "iol")
	}

	// Rule 3: Passwords must contain two different non-overlapping letter pairs.
	rule3 := func(password string) bool {
		m := make(map[rune]struct{})
		var lastR rune
		for _, r := range password {
			if r == lastR {
				m[r] = struct{}{}
			}
			lastR = r
		}
		return len(m) >= 2
	}

	return rule1(password) && rule2(password) && rule3(password)
}
