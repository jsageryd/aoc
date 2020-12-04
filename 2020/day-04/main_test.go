package main

import "testing"

func TestValid(t *testing.T) {
	for n, tc := range []struct {
		passport map[string]string
		valid    bool
	}{
		{map[string]string{"byr": "1937", "cid": "147", "ecl": "gry", "eyr": "2020", "hcl": "#fffffd", "hgt": "183cm", "iyr": "2017", "pid": "860033327"}, true},
		{map[string]string{"byr": "1929", "cid": "350", "ecl": "amb", "eyr": "2023", "hcl": "#cfa07d", "iyr": "2013", "pid": "028048884"}, false},
		{map[string]string{"byr": "1931", "ecl": "brn", "eyr": "2024", "hcl": "#ae17e1", "hgt": "179cm", "iyr": "2013", "pid": "760753108"}, true},
		{map[string]string{"ecl": "brn", "eyr": "2025", "hcl": "#cfa07d", "hgt": "59in", "iyr": "2011", "pid": "166559648"}, false},
	} {
		if got, want := valid(tc.passport), tc.valid; got != want {
			t.Errorf("[%d] got %t, want %t", n, got, want)
		}
	}
}
