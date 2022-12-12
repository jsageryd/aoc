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
	var v any
	json.Unmarshal(input, &v)

	var sum func(v any) int

	sum = func(v any) int {
		var s int

		switch val := v.(type) {
		case float64:
			s = int(val)
		case map[string]any:
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
		case []any:
			for _, e := range val {
				s += sum(e)
			}
		}

		return s
	}

	return sum(v)
}
