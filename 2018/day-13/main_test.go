package main

import (
	"bytes"
	"testing"
)

func TestRace(t *testing.T) {
	for n, tc := range []struct {
		tracks        []byte
		firstCollider *cart
		lastSurvivor  *cart
	}{
		{
			tracks: []byte(`
/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/`),
			firstCollider: &cart{id: 0, heading: north, x: 7, y: 3, nextTurn: left},
			lastSurvivor:  nil,
		},
		{
			tracks: []byte(`
/----/---\--------------------------------\
|    |   |                                |
\->--\---\-----------------------------<--/
     |   |
     \---/`),
			firstCollider: &cart{id: 0, heading: east, x: 13, y: 2, nextTurn: left},
			lastSurvivor:  nil,
		},
		{
			tracks: []byte(`
/>-<\
|   |
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`),
			firstCollider: &cart{id: 1, heading: west, x: 2, y: 0, nextTurn: left},
			lastSurvivor:  &cart{id: 8, heading: north, x: 6, y: 4, nextTurn: left},
		},
	} {
		tracks := bytes.Split(bytes.TrimSpace(tc.tracks), []byte("\n"))
		carts := extractCarts(tracks)

		firstCollider, lastSurvivor := race(tracks, carts)

		if *firstCollider != *tc.firstCollider {
			t.Errorf("[%d] firstCollider = (%s), want (%s)", n, firstCollider, tc.firstCollider)
		}

		switch {
		case lastSurvivor == nil && tc.lastSurvivor != nil:
			t.Errorf("[%d] lastSurvivor is nil, want (%s)", n, tc.lastSurvivor)
		case lastSurvivor != nil && tc.lastSurvivor == nil:
			t.Errorf("[%d] lastSurvivor = (%s), want nil", n, lastSurvivor)
		case lastSurvivor != nil && tc.lastSurvivor != nil:
			if *lastSurvivor != *tc.lastSurvivor {
				t.Errorf("[%d] lastSurvivor = (%s), want (%s)", n, lastSurvivor, tc.lastSurvivor)
			}
		}
	}
}

func TestExtractCarts(t *testing.T) {
	tracks := bytes.Split(bytes.TrimSpace([]byte(`
/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
  `)), []byte("\n"))

	wantTracks := bytes.Split(bytes.TrimSpace([]byte(`
/---\
|   |  /----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \-+--/
  \------/
  `)), []byte("\n"))

	wantCarts := []*cart{
		{id: 0, heading: east, x: 2, y: 0},
		{id: 1, heading: south, x: 9, y: 3},
	}

	carts := extractCarts(tracks)

	if got, want := len(tracks), len(wantTracks); got != want {
		t.Errorf("got %d track rows, want %d", got, want)
	} else {
		for n := range wantTracks {
			if !bytes.Equal(tracks[n], wantTracks[n]) {
				t.Errorf("tracks =\n%s\nwant\n%s", bytes.Join(tracks, []byte("\n")), bytes.Join(wantTracks, []byte("\n")))
				break
			}
		}
	}

	if len(carts) != len(wantCarts) {
		t.Errorf("carts = %q, want %q", carts, wantCarts)
	} else {
		for n := range wantCarts {
			if *carts[n] != *wantCarts[n] {
				t.Errorf("carts = %q, want %q", carts, wantCarts)
				break
			}
		}
	}
}
