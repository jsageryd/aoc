package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		totalLength         int
		totalLengthUnquoted int
		totalLengthRequoted int
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		totalLength += len(scanner.Text())
		totalLengthUnquoted += len(unquote(scanner.Text()))
		totalLengthRequoted += len(quote(scanner.Text()))
	}

	fmt.Printf("Part 1: %d\n", totalLength-totalLengthUnquoted)
	fmt.Printf("Part 2: %d\n", totalLengthRequoted-totalLength)
}

func unquote(s string) string {
	s, _ = strconv.Unquote(s)
	return s
}

func quote(s string) string {
	return strconv.Quote(s)
}
