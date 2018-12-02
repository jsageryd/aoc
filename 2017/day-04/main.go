package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var validPassphrases int
	var validPassphrasesWithAnagramCheck int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if valid(scanner.Text()) {
			validPassphrases++
		}
		if validWithAnagramCheck(scanner.Text()) {
			validPassphrasesWithAnagramCheck++
		}
	}

	fmt.Printf("Part 1: %d\n", validPassphrases)
	fmt.Printf("Part 2: %d\n", validPassphrasesWithAnagramCheck)
}

func valid(passphrase string) bool {
	words := strings.Split(passphrase, " ")

	seen := map[string]bool{}

	for _, w := range words {
		if _, ok := seen[w]; ok {
			return false
		}
		seen[w] = true
	}

	return true
}

func validWithAnagramCheck(passphrase string) bool {
	words := strings.Split(passphrase, " ")

	seen := map[string]bool{}

	for _, w := range words {
		if _, ok := seen[letterOrderIndependentHash(w)]; ok {
			return false
		}
		seen[letterOrderIndependentHash(w)] = true
	}

	return true
}

func letterOrderIndependentHash(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	var res []rune
	count := 0
	for n := range runes {
		count++
		if n >= len(runes)-1 || runes[n] != runes[n+1] {
			res = append(res, runes[n])
			res = append(res, []rune(strconv.Itoa(count))...)
			count = 0
		}
	}
	return string(res)
}
