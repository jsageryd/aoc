package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := []string{
		"Hit Points: 1",
		"Damage: 2",
		"Armor: 3",
	}

	gotC := parse(input)

	wantC := contender{hitPoints: 1, damage: 2, armor: 3}

	if gotC != wantC {
		t.Errorf("got %+v, want %+v", gotC, wantC)
	}
}

func TestPlayerWins(t *testing.T) {
	player := contender{hitPoints: 8, damage: 5, armor: 5}
	enemy := contender{hitPoints: 12, damage: 7, armor: 2}

	if got, want := playerWins(&player, &enemy), true; got != want {
		t.Errorf("got %t, want %t", got, want)
	}

	if got, want := player.hitPoints, 2; got != want {
		t.Errorf("player has %d hit points, want %d", got, want)
	}

	if got, want := enemy.hitPoints, 0; got != want {
		t.Errorf("enemy has %d hit points, want %d", got, want)
	}
}
