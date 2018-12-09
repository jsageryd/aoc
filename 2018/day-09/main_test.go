package main

import "testing"

func TestPlay(t *testing.T) {
	for n, tc := range []struct {
		players      int
		lastMarble   int
		highestScore int
	}{
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	} {
		highestScore := play(tc.players, tc.lastMarble)

		if highestScore != tc.highestScore {
			t.Errorf("[%d] play(%d, %d) = %d, want %d",
				n, tc.players, tc.lastMarble,
				highestScore, tc.highestScore,
			)
		}
	}
}

func TestCircle(t *testing.T) {
	c := &circle{}

	for _, tc := range []struct {
		desc       string
		op         func()
		wantData   []int
		wantCurIdx int
	}{
		{
			desc:       "place0",
			op:         func() { c.place(0) },
			wantData:   []int{0},
			wantCurIdx: 0,
		},
		{
			desc:       "place1",
			op:         func() { c.place(1) },
			wantData:   []int{0, 1},
			wantCurIdx: 1,
		},
		{
			desc:       "place2",
			op:         func() { c.place(2) },
			wantData:   []int{0, 2, 1},
			wantCurIdx: 1,
		},
		{
			desc:       "place3",
			op:         func() { c.place(3) },
			wantData:   []int{0, 2, 1, 3},
			wantCurIdx: 3,
		},
		{
			desc:       "place4",
			op:         func() { c.place(4) },
			wantData:   []int{0, 4, 2, 1, 3},
			wantCurIdx: 1,
		},
		{
			desc:       "place5",
			op:         func() { c.place(5) },
			wantData:   []int{0, 4, 2, 5, 1, 3},
			wantCurIdx: 3,
		},
		{
			desc:       "place6",
			op:         func() { c.place(6) },
			wantData:   []int{0, 4, 2, 5, 1, 6, 3},
			wantCurIdx: 5,
		},
		{
			desc:       "delete",
			op:         func() { c.delete(0) },
			wantData:   []int{0, 4, 2, 5, 1, 3},
			wantCurIdx: 5,
		},
		{
			desc:       "deleteNegativeOffset",
			op:         func() { c.delete(-7) },
			wantData:   []int{0, 4, 2, 5, 3},
			wantCurIdx: 4,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			t.Logf("before: %s", c)
			tc.op()
			t.Logf(" after: %s", c)

			if got, want := c.curIdx, tc.wantCurIdx; got != want {
				t.Errorf("curIdx = %d, want %d", got, want)
			}

			if got, want := c.data, tc.wantData; !intSlicesEqual(got, want) {
				t.Errorf("data = %v, want %v", got, want)
			}
		})
	}
}

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for n := range a {
		if a[n] != b[n] {
			return false
		}
	}
	return true
}
