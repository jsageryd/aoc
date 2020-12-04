package main

import "testing"

func TestValid(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		for n, tc := range []struct {
			passport map[string]string
			valid    bool
		}{
			{map[string]string{"byr": "1937", "cid": "147", "ecl": "gry", "eyr": "2020", "hcl": "#fffffd", "hgt": "183cm", "iyr": "2017", "pid": "860033327"}, true},
			{map[string]string{"byr": "1929", "cid": "350", "ecl": "amb", "eyr": "2023", "hcl": "#cfa07d", "iyr": "2013", "pid": "028048884"}, false},
			{map[string]string{"byr": "1931", "ecl": "brn", "eyr": "2024", "hcl": "#ae17e1", "hgt": "179cm", "iyr": "2013", "pid": "760753108"}, true},
			{map[string]string{"ecl": "brn", "eyr": "2025", "hcl": "#cfa07d", "hgt": "59in", "iyr": "2011", "pid": "166559648"}, false},
		} {
			if got, want := valid(tc.passport, false), tc.valid; got != want {
				t.Errorf("[%d] got %t, want %t", n, got, want)
			}
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		for n, tc := range []struct {
			passport map[string]string
			valid    bool
		}{
			{map[string]string{"byr": "1926", "cid": "100", "ecl": "amb", "eyr": "1972", "hcl": "#18171d", "hgt": "170", "iyr": "2018", "pid": "186cm"}, false},
			{map[string]string{"byr": "1946", "ecl": "grn", "eyr": "1967", "hcl": "#602927", "hgt": "170cm", "iyr": "2019", "pid": "012533040"}, false},
			{map[string]string{"byr": "1992", "cid": "277", "ecl": "brn", "eyr": "2020", "hcl": "dab227", "hgt": "182cm", "iyr": "2012", "pid": "021572410"}, false},
			{map[string]string{"byr": "2007", "ecl": "zzz", "eyr": "2038", "hcl": "74454a", "hgt": "59cm", "iyr": "2023", "pid": "3556412378"}, false},
			{map[string]string{"byr": "1980", "ecl": "grn", "eyr": "2030", "hcl": "#623a2f", "hgt": "74in", "iyr": "2012", "pid": "087499704"}, true},
			{map[string]string{"byr": "1989", "cid": "129", "ecl": "blu", "eyr": "2029", "hcl": "#a97842", "hgt": "165cm", "iyr": "2014", "pid": "896056539"}, true},
			{map[string]string{"hcl": "#888785", "hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88", "pid": "545766238", "ecl": "hzl", "eyr": "2022"}, true},
			{map[string]string{"iyr": "2010", "hgt": "158cm", "hcl": "#b6652a", "ecl": "blu", "byr": "1944", "eyr": "2021", "pid": "093154719"}, true},
		} {
			if got, want := valid(tc.passport, true), tc.valid; got != want {
				t.Errorf("[%d] got %t, want %t", n, got, want)
			}
		}
	})
}

func TestValidField(t *testing.T) {
	for n, tc := range []struct {
		field, value string
		valid        bool
	}{
		{"byr", "1920", true},
		{"byr", "2002", true},
		{"byr", "1919", false},
		{"byr", "2003", false},

		{"iyr", "2010", true},
		{"iyr", "2020", true},
		{"iyr", "2009", false},
		{"iyr", "2021", false},

		{"eyr", "2020", true},
		{"eyr", "2030", true},
		{"eyr", "2019", false},
		{"eyr", "2031", false},

		{"hgt", "150cm", true},
		{"hgt", "193cm", true},
		{"hgt", "149cm", false},
		{"hgt", "194cm", false},
		{"hgt", "60in", true},
		{"hgt", "76in", true},
		{"hgt", "58in", false},
		{"hgt", "77in", false},
		{"hgt", "190", false},

		{"hcl", "#123abc", true},
		{"hcl", "#123abz", false},
		{"hcl", "123abc", false},

		{"ecl", "amb", true},
		{"ecl", "blu", true},
		{"ecl", "brn", true},
		{"ecl", "gry", true},
		{"ecl", "grn", true},
		{"ecl", "hzl", true},
		{"ecl", "oth", true},
		{"ecl", "foo", false},

		{"pid", "000000001", true},
		{"pid", "01234567", false},
		{"pid", "0123456789", false},
	} {
		if got, want := validField(tc.field, tc.value), tc.valid; got != want {
			t.Errorf("[%d] validField(%q, %q) = %t, want %t", n, tc.field, tc.value, got, want)
		}
	}
}
