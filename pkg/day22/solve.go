package day22

import (
	"adventofcode2015/pkg/utils"
	"math"
)

func Solve(partTwo bool) int {
	game := NewGame(50, 500, 51, 9, partTwo)
	return CheapestWinningSpell(game)
}

var Cache = make(map[Game]int)

func Memoize(game Game, fn func(game Game) int) int {
	value, exists := Cache[game]
	if exists {
		return value
	}
	value = fn(game)
	Cache[game] = value
	return value
}

func CheapestWinningSpell(initial Game) int {
	return Memoize(initial, func(initial Game) int {
		best := math.MaxInt
		for _, spell := range Spells {
			// fmt.Printf("%v. %v (best: %v)\n", current, spell.name, best)
			game := initial

			switch game.Turn(spell) {
			case PlayerWon:
				best = utils.MinValue(best, spell.mana)
			case PlayerWonWithoutCasting:
				// cases where poison was applied before we cast anything
				// resulting in the boss dying
				best = 0
			case TurnComplete:
				// ensure's that _something_ could be cast winning the game
				if candidate := CheapestWinningSpell(game); candidate != math.MaxInt {
					best = utils.MinValue(best, spell.mana+candidate)
				}
			default:
				// ignore BossWon, OutOfMana, or EffectExists as we can
				// just let the loop continue and try the next spell
			}
		}

		return best
	})
}
