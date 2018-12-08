package main

import "testing"

func TestFloor(t *testing.T) {
	for n, tc := range []struct {
		input []byte
		floor int
	}{
		{[]byte("(())"), 0},
		{[]byte("()()"), 0},
		{[]byte("((("), 3},
		{[]byte("(()(()("), 3},
		{[]byte("))((((("), 3},
		{[]byte("())"), -1},
		{[]byte("))("), -1},
		{[]byte(")))"), -3},
		{[]byte(")())())"), -3},
	} {
		if got, want := floor(tc.input), tc.floor; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestFirstBasementPosition(t *testing.T) {
	for n, tc := range []struct {
		input    []byte
		position int
	}{
		{[]byte(")"), 1},
		{[]byte("()())"), 5},
	} {
		if got, want := firstBasementPosition(tc.input), tc.position; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
