package main

import "testing"

func TestTotalScore(t *testing.T) {
	for n, tc := range []struct {
		input   []byte
		score   int
		garbage int
	}{
		{[]byte(`{}`), 1, 0},
		{[]byte(`{{{}}}`), 6, 0},
		{[]byte(`{{},{}}`), 5, 0},
		{[]byte(`{{{},{},{{}}}}`), 16, 0},
		{[]byte(`{<a>,<a>,<a>,<a>}`), 1, 4},
		{[]byte(`{{<ab>},{<ab>},{<ab>},{<ab>}}`), 9, 8},
		{[]byte(`{{<!!>},{<!!>},{<!!>},{<!!>}}`), 9, 0},
		{[]byte(`{{<a!>},{<a!>},{<a!>},{<ab>}}`), 3, 17},
		{[]byte(`<>`), 0, 0},
		{[]byte(`<random characters>`), 0, 17},
		{[]byte(`<<<<>`), 0, 3},
		{[]byte(`<{!>}>`), 0, 2},
		{[]byte(`<!!>`), 0, 0},
		{[]byte(`<!!!>>`), 0, 0},
		{[]byte(`<{o"i!a,<{i<a>`), 0, 10},
	} {
		score, garbage := scoreAndGarbage(tc.input)

		if score != tc.score || garbage != tc.garbage {
			t.Errorf("[%d] got %d, %d; want %d, %d", n, score, garbage, tc.score, tc.garbage)
		}
	}
}
