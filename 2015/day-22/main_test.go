package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte(`
Hit Points: 1
Damage: 2
`)[1:]

	gotE := parse(input)

	wantE := enemy{hitPoints: 1, damage: 2}

	if gotE != wantE {
		t.Errorf("got %+v, want %+v", gotE, wantE)
	}
}

func TestPerm(t *testing.T) {
	for n, tc := range []struct {
		n        int
		wantSeqs [][]int
	}{
		{
			n:        0,
			wantSeqs: [][]int{},
		},
		{
			n: 1,
			wantSeqs: [][]int{
				{0}, {0, 0}, {0, 0, 0},
			},
		},
		{
			n: 2,
			wantSeqs: [][]int{
				{0}, {0, 0}, {0, 0, 0}, {0, 0, 1}, {0, 1}, {0, 1, 0}, {0, 1, 1},
				{1}, {1, 0}, {1, 0, 0}, {1, 0, 1}, {1, 1}, {1, 1, 0}, {1, 1, 1},
			},
		},
	} {
		var gotSeqs [][]int

		perm(tc.n, func(seq []int) bool {
			s := make([]int, len(seq))
			copy(s, seq)
			gotSeqs = append(gotSeqs, s)
			return len(seq) < 3
		})

		if fmt.Sprint(gotSeqs) != fmt.Sprint(tc.wantSeqs) {
			t.Errorf("[%d] got %d, want %d", n, gotSeqs, tc.wantSeqs)
		}
	}
}

func TestPlayerWins(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		for _, tc := range []struct {
			desc                string
			player              player
			enemy               enemy
			spellIdxs           []int
			hardMode            bool
			wantWin             bool
			wantPlayerHitPoints int
			wantEnemyHitPoints  int
		}{
			{
				desc:                "First example from description",
				player:              player{hitPoints: 10, mana: 250},
				enemy:               enemy{hitPoints: 13, damage: 8},
				spellIdxs:           []int{3, 0},
				hardMode:            false,
				wantWin:             true,
				wantPlayerHitPoints: 2,
				wantEnemyHitPoints:  0,
			},
			{
				desc:                "Second example from description",
				player:              player{hitPoints: 10, mana: 250},
				enemy:               enemy{hitPoints: 14, damage: 8},
				spellIdxs:           []int{4, 2, 1, 3, 0},
				hardMode:            false,
				wantWin:             true,
				wantPlayerHitPoints: 1,
				wantEnemyHitPoints:  -1,
			},
			{
				desc:                "Enemy wins",
				player:              player{hitPoints: 10, mana: 500},
				enemy:               enemy{hitPoints: 13, damage: 8},
				spellIdxs:           []int{4, 3},
				hardMode:            false,
				wantWin:             false,
				wantPlayerHitPoints: -6,
				wantEnemyHitPoints:  10,
			},
			{
				desc:                "Insufficient mana",
				player:              player{hitPoints: 10, mana: 10},
				enemy:               enemy{hitPoints: 13, damage: 8},
				spellIdxs:           []int{0},
				hardMode:            false,
				wantWin:             false,
				wantPlayerHitPoints: 10,
				wantEnemyHitPoints:  13,
			},
			{
				desc:                "Damage is always at least 1",
				player:              player{hitPoints: 10, mana: 500},
				enemy:               enemy{hitPoints: 5, damage: 7},
				spellIdxs:           []int{2, 0, 0},
				hardMode:            false,
				wantWin:             true,
				wantPlayerHitPoints: 8,
				wantEnemyHitPoints:  -3,
			},
			{
				desc:                "First example from description with hard mode",
				player:              player{hitPoints: 10, mana: 250},
				enemy:               enemy{hitPoints: 13, damage: 8},
				spellIdxs:           []int{3, 0},
				hardMode:            true,
				wantWin:             false,
				wantPlayerHitPoints: 0,
				wantEnemyHitPoints:  7,
			},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				p, e := tc.player, tc.enemy

				gotWin, err := playerWins(&p, &e, tc.spellIdxs, tc.hardMode)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if gotWin != tc.wantWin {
					t.Errorf("got %t, want %t", gotWin, tc.wantWin)
				}

				if got, want := p.hitPoints, tc.wantPlayerHitPoints; got != want {
					t.Errorf("player has %d hit points, want %d", got, want)
				}

				if got, want := e.hitPoints, tc.wantEnemyHitPoints; got != want {
					t.Errorf("enemy has %d hit points, want %d", got, want)
				}
			})
		}
	})

	t.Run("Errors", func(t *testing.T) {
		for _, tc := range []struct {
			desc      string
			player    player
			enemy     enemy
			spellIdxs []int
			wantErr   error
		}{
			{
				desc:      "Not enough spells",
				player:    player{hitPoints: 10, mana: 500},
				enemy:     enemy{hitPoints: 13, damage: 8},
				spellIdxs: []int{0},
				wantErr:   errNoMoreSpells,
			},
			{
				desc:      "Invalid spell combination",
				player:    player{hitPoints: 10, mana: 500},
				enemy:     enemy{hitPoints: 13, damage: 8},
				spellIdxs: []int{2, 2},
				wantErr:   errInvalidSpellCombination,
			},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				p, e := tc.player, tc.enemy

				gotWin, gotErr := playerWins(&p, &e, tc.spellIdxs, false)

				if got, want := gotWin, false; got != want {
					t.Errorf("got %t, want %t", got, want)
				}

				if gotErr == nil {
					t.Fatal("error is nil")
				}

				if gotErr != tc.wantErr {
					t.Errorf("error is %q, want %q", gotErr, tc.wantErr)
				}
			})
		}
	})
}
