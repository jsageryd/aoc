package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", sumOfNumbers(input))
}

func sumOfNumbers(input []byte) int {
	var v interface{}
	json.Unmarshal(input, &v)

	var sum func(v interface{}) int

	sum = func(v interface{}) int {
		var s int

		switch val := v.(type) {
		case float64:
			s = int(val)
		case map[string]interface{}:
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
