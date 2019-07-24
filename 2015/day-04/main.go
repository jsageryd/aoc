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

	n := int64(1)

	for !md5HexHasFiveLeadingZeros(b) {
		n++
		b = strconv.AppendInt(b[:len(input)], n, 10)
	}

	fmt.Printf("Part 1: %d\n", n)
}

func md5HexHasFiveLeadingZeros(b []byte) bool {
	s := md5.Sum(b)
	return strings.HasPrefix(hex.EncodeToString(s[:]), "00000")
}
