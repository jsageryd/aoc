package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)

	code := parse(input)

	// Restore as per instructions
	code[1], code[2] = 12, 2

	result, err := run(code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", result[0])
}

func parse(code string) []int {
	ss := strings.Split(code, ",")
	ii := make([]int, 0, len(ss))

	for n := range ss {
		if i, err := strconv.Atoi(ss[n]); err == nil {
			ii = append(ii, i)
		}
	}

	return ii
}

func run(code []int) ([]int, error) {
	for cur := 0; cur < len(code); cur += 4 {
		switch op := code[cur]; op {
		case 1:
			code[code[cur+3]] = code[code[cur+1]] + code[code[cur+2]]
		case 2:
			code[code[cur+3]] = code[code[cur+1]] * code[code[cur+2]]
		case 99:
			return code, nil
		default:
			return nil, fmt.Errorf("position %d has value %d; want 1, 2 or 99", cur, op)
		}
	}

	return code, nil
}
