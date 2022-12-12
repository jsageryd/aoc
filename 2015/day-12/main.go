package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", sumOfNumbers(input, false))
	fmt.Printf("Part 2: %d\n", sumOfNumbers(input, true))
}

func sumOfNumbers(input []byte, skipRedObjects bool) int {
	var v interface{}
	json.Unmarshal(input, &v)

	var sum func(v interface{}) int

	sum = func(v interface{}) int {
		var s int

		switch val := v.(type) {
		case float64:
			s = int(val)
		case map[string]interface{}:
			if skipRedObjects {
				for _, v := range val {
					if s, ok := v.(string); ok && s == "red" {
						return 0
					}
				}
			}
			for _, v := range val {
				s += sum(v)
			}
		case []interface{}:
			for _, e := range val {
				s += sum(e)
			}
		}

		return s
	}

	return sum(v)
}
