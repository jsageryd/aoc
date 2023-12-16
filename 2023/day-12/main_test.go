package main

import (
	"fmt"
	"testing"
	"time"
)

var input = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 21; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		if got, want := part2(input), 525152; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}()

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		t.Error("test timed out")
	}
}

func TestArrangements(t *testing.T) {
	for n, tc := range []struct {
		row  string
		want int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	} {
		if got, want := arrangements(tc.row), tc.want; got != want {
			t.Errorf("[%d] arrangements(%q) = %d, want %d", n, tc.row, got, want)
		}
	}
}

func TestPerm(t *testing.T) {
	perms := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'a', 'b', 'd', 'c'},
		{'a', 'c', 'b', 'd'},
		{'a', 'c', 'd', 'b'},
		{'a', 'd', 'b', 'c'},
		{'a', 'd', 'c', 'b'},
		{'b', 'a', 'c', 'd'},
		{'b', 'a', 'd', 'c'},
		{'b', 'c', 'a', 'd'},
		{'b', 'c', 'd', 'a'},
		{'b', 'd', 'a', 'c'},
		{'b', 'd', 'c', 'a'},
		{'c', 'a', 'b', 'd'},
		{'c', 'a', 'd', 'b'},
		{'c', 'b', 'a', 'd'},
		{'c', 'b', 'd', 'a'},
		{'c', 'd', 'a', 'b'},
		{'c', 'd', 'b', 'a'},
		{'d', 'a', 'b', 'c'},
		{'d', 'a', 'c', 'b'},
		{'d', 'b', 'a', 'c'},
		{'d', 'b', 'c', 'a'},
		{'d', 'c', 'a', 'b'},
		{'d', 'c', 'b', 'a'},
	}

	for i := 0; i < len(perms)-2; i++ {
		if ok := perm(perms[i]); !ok {
			t.Errorf("[%d] got non-ok return", i)
		}
		if got, want := perms[i], perms[i+1]; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %q, want %q", i, got, want)
		}
	}

	if perm(perms[len(perms)-1]) {
		t.Errorf("got ok return value for last perm")
	}
}

func TestRev(t *testing.T) {
	for i, tc := range []struct {
		have, want []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte("a"), []byte("a")},
		{[]byte("ab"), []byte("ba")},
		{[]byte("abc"), []byte("cba")},
		{[]byte("abcd"), []byte("dcba")},
		{[]byte("abcde"), []byte("edcba")},
	} {
		got := make([]byte, len(tc.have))
		copy(got, tc.have)
		rev(got)
		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("[%d] got %q, want %q", i, got, tc.want)
		}
	}
}

func TestExpand(t *testing.T) {
	for n, tc := range []struct {
		row  string
		want string
	}{
		{row: ".# 1", want: ".#?.#?.#?.#?.# 1,1,1,1,1"},
		{row: "???.### 1,1,3", want: "???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3"},
	} {
		if got, want := expand(tc.row), tc.want; got != want {
			t.Errorf("[%d] expand(%q) = %q, want %q", n, tc.row, got, want)
		}
	}
}
