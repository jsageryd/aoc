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
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		totalLength += len(scanner.Text())
		totalLengthUnquoted += len(unquote(scanner.Text()))
	}

	fmt.Printf("Part 1: %d\n", totalLength-totalLengthUnquoted)
}

func unquote(s string) string {
	s, _ = strconv.Unquote(s)
	return s
}
