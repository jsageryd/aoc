package main

import (
	"errors"
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

	result2, err := runUntil(parse(input), 19690720)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", result[0])
	fmt.Printf("Part 2: %d\n", 100*result2[1]+result2[2])
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

func runUntil(code []int, firstValue int) ([]int, error) {
	c := make([]int, len(code))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(c, code)
			c[1], c[2] = noun, verb
			c, err := run(c)
			if err != nil {
				return nil, fmt.Errorf("[noun %d, verb %d] %v", noun, verb, err)
			}
			if c[0] == firstValue {
				return c, nil
			}
		}
	}

	return nil, errors.New("target value not found")
}
