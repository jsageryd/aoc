package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var passports []map[string]string

	passport := map[string]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		for _, element := range strings.Fields(row) {
			kv := strings.Split(element, ":")
			passport[kv[0]] = kv[1]
		}
	}
	passports = append(passports, passport)

	var validCount int

	for _, p := range passports {
		if valid(p) {
			validCount++
		}
	}

	fmt.Printf("Part 1: %d\n", validCount)
}

func valid(passport map[string]string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, f := range requiredFields {
		if _, ok := passport[f]; !ok {
			return false
		}
	}

	return true
}
