package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	var validCountPart1 int
	var validCountPart2 int

	for _, p := range passports {
		if valid(p, false) {
			validCountPart1++
		}
		if valid(p, true) {
			validCountPart2++
		}
	}

	fmt.Printf("Part 1: %d\n", validCountPart1)
	fmt.Printf("Part 2: %d\n", validCountPart2)
}

func valid(passport map[string]string, validateFieldData bool) bool {
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
		v, ok := passport[f]
		if !ok {
			return false
		}
		if validateFieldData && !validField(f, v) {
			return false
		}
	}

	return true
}

var validHCL = regexp.MustCompile("^#[0-9a-f]{6}$")
var validPID = regexp.MustCompile("^[0-9]{9}$")

func validField(field, value string) bool {
	switch field {
	case "byr":
		yearInt, err := strconv.Atoi(value)
		return err == nil && yearInt >= 1920 && yearInt <= 2002
	case "iyr":
		yearInt, err := strconv.Atoi(value)
		return err == nil && yearInt >= 2010 && yearInt <= 2020
	case "eyr":
		yearInt, err := strconv.Atoi(value)
		return err == nil && yearInt >= 2020 && yearInt <= 2030
	case "hgt":
		switch {
		case strings.HasSuffix(value, "cm"):
			lengthInt, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
			return err == nil && lengthInt >= 150 && lengthInt <= 193
		case strings.HasSuffix(value, "in"):
			lengthInt, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
			return err == nil && lengthInt >= 59 && lengthInt <= 76
		default:
			return false
		}
	case "hcl":
		return validHCL.MatchString(value)
	case "ecl":
		switch value {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	case "pid":
		return validPID.MatchString(value)
	default:
		return true
	}
}
