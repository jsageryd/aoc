package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	input = bytes.TrimSpace(input)

	var n, n1, n2 int

	for n1 == 0 || n2 == 0 {
		n++

		s := md5.Sum(strconv.AppendInt(input[:], int64(n), 10))
		z := leadingHexZeros(s[:])

		if n1 == 0 && z >= 5 {
			n1 = n
		}

		if n2 == 0 && z >= 6 {
			n2 = n
		}
	}

	fmt.Printf("Part 1: %d\n", n1)
	fmt.Printf("Part 2: %d\n", n2)
}

func leadingHexZeros(input []byte) int {
	var c int
loop:
	for n := 0; n < len(input); n++ {
		switch {
		case input[n] == 0:
			c += 2
		case input[n] <= 0x0f:
			c++
		default:
			break loop
		}
	}
	return c
}
