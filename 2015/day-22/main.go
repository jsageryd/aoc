package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	fmt.Printf("Part 1: %d\n", run(input, false))
	fmt.Printf("Part 2: %d\n", run(input, true))
}

func run(input []byte, hardMode bool) int {
	p := player{hitPoints: 50, mana: 500}
	e := parse(input)

	minMana := 10000
	maxSpells := 10000

	perm(len(spells), func(spellIndices []int) bool {
		if len(spellIndices) > maxSpells {
			return false
		}

		debugLogf("Trying spells %d", spellIndices)

		p, e := p, e

		win, err := playerWins(&p, &e, spellIndices, hardMode)
		if err != nil {
			switch err {
			case errNoMoreSpells:
				return true // try again with longer list of spells
			case errInvalidSpellCombination:
				return false // branch not valid
			default:
				fmt.Fprintf(os.Stderr, "Unknown error: %v\n", err)
				os.Exit(1)
			}
		}

		// If player wins, check spent mana and stop branch
		if win {
			debugLogf("Win with spells %d and %d mana spent", spellIndices, p.manaSpent)

			if p.manaSpent < minMana {
				minMana = p.manaSpent
				maxSpells = len(spellIndices)
			}

			return false
		}

		// If player loses, he cannot win by adding more spells
		return false
	})

	return minMana
}

var debugLogf = func(format string, v ...any) {}

func init() {
	if debug, _ := strconv.ParseBool(os.Getenv("DEBUG")); debug {
		debugLogf = log.New(os.Stderr, "", 0).Printf
	}
}

type spell struct {
	name   string
	cost   int
	timer  int
	damage int
	heal   int
	armor  int
	mana   int
}

var spells = []spell{
	{name: "Magic Missile", cost: 53, damage: 4},
	{name: "Drain", cost: 73, damage: 2, heal: 2},
	{name: "Shield", cost: 113, timer: 6, armor: 7},
	{name: "Poison", cost: 173, timer: 6, damage: 3},
	{name: "Recharge", cost: 229, timer: 5, mana: 101},
}

func parse(input []byte) enemy {
	var e enemy
	fmt.Sscanf(string(input), "Hit Points: %d\nDamage: %d\n", &e.hitPoints, &e.damage)
	return e
}

var (
	errInvalidSpellCombination = errors.New("invalid spell combination")
	errNoMoreSpells            = errors.New("no more spells")
)

// perm generates permutations of numbers 0 to n, calling f for each sequence.
// perm continues until f returns false for all branches.
func perm(n int, f func(seq []int) bool) {
	var rec func(idx int, seq []int)

	rec = func(idx int, seq []int) {
		if len(seq) < idx+1 {
			seq = append(seq, 0)
		}

		for i := 0; i < n; i++ {
			seq[idx] = i
			if f(seq) {
				rec(idx+1, seq)
			}
		}
	}

	rec(0, nil)
}

type player struct {
	hitPoints int
	armor     int
	mana      int
	manaSpent int
}

type enemy struct {
	hitPoints int
	damage    int
}

// playerWins runs the fight and returns true if player wins.
func playerWins(p *player, e *enemy, spellIndices []int, hardMode bool) (win bool, err error) {
	defer func() {
		if win {
			debugLogf("Player wins")
		} else {
			debugLogf("Enemy wins")
		}
		debugLogf("Player: %+v", *p)
		debugLogf("Enemy: %+v", *e)
		debugLogf("")
	}()

	var spellList []spell

	for _, idx := range spellIndices {
		spellList = append(spellList, spells[idx])
	}

	apply := func(s *spell) {
		e.hitPoints -= s.damage
		p.hitPoints += s.heal
		p.armor += s.armor
		p.mana += s.mana
		if s.timer > 0 {
			s.timer--
		}
		debugLogf("Apply %s; timer is %d", s.name, s.timer)
	}

	effects := make(map[string]*spell)

	for turn := 0; ; turn++ {
		debugLogf("Turn %d", turn)
		p.armor = 0
		for _, s := range effects {
			apply(s)
			if s.timer <= 0 {
				delete(effects, s.name)
			}
		}

		if e.hitPoints <= 0 {
			debugLogf("Enemy dies")
			return true, nil
		}

		playersTurn := turn%2 == 0

		if playersTurn {
			if hardMode {
				p.hitPoints--

				if p.hitPoints <= 0 {
					debugLogf("Player dies because game too hard")
					return false, nil
				}
			}

			if len(spellList) == 0 {
				debugLogf("Player runs out of spells")
				return false, errNoMoreSpells
			}

			var s spell
			s, spellList = spellList[0], spellList[1:]

			if p.mana < s.cost {
				debugLogf("Player runs out of mana")
				return false, nil
			}

			p.mana -= s.cost
			p.manaSpent += s.cost

			debugLogf("Player casts %s", s.name)

			if s.timer > 0 {
				if _, ok := effects[s.name]; ok {
					debugLogf("Invalid spell combination")
					return false, errInvalidSpellCombination
				}
				effects[s.name] = &s
			} else {
				apply(&s)
			}
		} else {
			damage := e.damage - p.armor
			if damage < 1 {
				damage = 1
			}
			debugLogf("Enemy deals %d damage", damage)
			p.hitPoints -= damage
		}

		if e.hitPoints <= 0 {
			debugLogf("Enemy dies")
			return true, nil
		}

		if p.hitPoints <= 0 {
			debugLogf("Player dies")
			return false, nil
		}

		debugLogf("Player: %+v", *p)
		debugLogf("Enemy: %+v", *e)
		debugLogf("")
	}

	return false, errors.New("loop not infinite")
}
