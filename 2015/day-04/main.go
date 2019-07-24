package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	input = bytes.TrimSpace(input)

	b := append(input, '1')

	var n1, n2 int64
	n := int64(1)

	for n1 == 0 || n2 == 0 {
		s := md5.Sum(b)

		if n1 == 0 && strings.HasPrefix(hex.EncodeToString(s[:]), "00000") {
			n1 = n
		}

		if n2 == 0 && strings.HasPrefix(hex.EncodeToString(s[:]), "000000") {
			n2 = n
		}

		n++
		b = strconv.AppendInt(b[:len(input)], n, 10)
	}

	fmt.Printf("Part 1: %d\n", n1)
	fmt.Printf("Part 2: %d\n", n2)
}
