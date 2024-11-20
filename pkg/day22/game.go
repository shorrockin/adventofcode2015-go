package day22

import (
	"adventofcode2015/pkg/utils"
)

type Game struct {
	playerHitpoints   int
	playerMana        int
	bossHitpoints     int
	bossDamage        int
	shieldTurnsLeft   int
	poisonTurnsLeft   int
	rechargeTurnsLeft int
	partTwo           bool
}

func NewGame(playerHitpoints, playerMana, bossHitpoints, bossDamage int, partTwo bool) Game {
	return Game{
		playerHitpoints:   playerHitpoints,
		playerMana:        playerMana,
		bossHitpoints:     bossHitpoints,
		bossDamage:        bossDamage,
		shieldTurnsLeft:   0,
		poisonTurnsLeft:   0,
		rechargeTurnsLeft: 0,
		partTwo:           partTwo,
	}
}

type TurnResult int

const (
	PlayerWon TurnResult = iota
	PlayerWonWithoutCasting
	BossWon
	OutOfMana
	EffectExists
	TurnComplete
)

func (game *Game) Turn(spell Spell) TurnResult {
	if spell.mana > game.playerMana {
		return OutOfMana
	}
	if spell.shield > 0 && game.shieldTurnsLeft > 1 {
		return EffectExists
	}
	if spell.poison > 0 && game.poisonTurnsLeft > 1 {
		return EffectExists
	}
	if spell.recharge > 0 && game.rechargeTurnsLeft > 1 {
		return EffectExists
	}

	if game.partTwo {
		game.playerHitpoints -= 1
		if game.playerHitpoints <= 0 {
			return BossWon
		}
	}

	//fmt.Println("\n-- Player turn --")
	game.applyEffects()
	if game.bossHitpoints <= 0 {
		return PlayerWonWithoutCasting
	}

	//fmt.Printf("Player casts %v dealing %v damage\n", spell.name, spell.damage)
	game.playerMana -= spell.mana
	game.playerHitpoints += spell.heal
	game.bossHitpoints -= spell.damage
	game.shieldTurnsLeft += spell.shield
	game.poisonTurnsLeft += spell.poison
	game.rechargeTurnsLeft += spell.recharge

	armor := 0
	if game.shieldTurnsLeft > 0 {
		armor = 7
	}

	//fmt.Println("\n-- Boss turn --")
	game.applyEffects()
	//fmt.Printf("Boss attacks for 8 - %v = %v damage\n", armor, utils.MaxValue(1, 8-armor))
	game.playerHitpoints -= utils.MaxValue(1, 8-armor)
	if game.playerHitpoints <= 0 {
		return BossWon
	}

	if game.bossHitpoints <= 0 {
		return PlayerWon
	}

	return TurnComplete
}

func (game *Game) applyEffects() {
	// armor := 0
	// if game.shieldTurnsLeft > 0 {
	// 	armor = 7
	// }
	//fmt.Printf("- Player has %v hit points, %v armor, %v mana\n", game.playerHitpoints, armor, game.playerMana)
	//fmt.Printf("- Boss has %v hit points\n", game.bossHitpoints)
	if game.shieldTurnsLeft > 0 {
		game.shieldTurnsLeft -= 1
		//fmt.Printf("Shield's timer is now %v\n", game.shieldTurnsLeft)
	}
	if game.poisonTurnsLeft > 0 {
		game.poisonTurnsLeft -= 1
		game.bossHitpoints -= 3
		//fmt.Printf("Poison deals 3 damage; it's timer is now %v\n", game.poisonTurnsLeft)
	}
	if game.rechargeTurnsLeft > 0 {
		game.rechargeTurnsLeft -= 1
		game.playerMana += 101
		//fmt.Printf("Rechange provided 101 mana; it's time is now %v\n", game.rechargeTurnsLeft)
	}
}

type Spell struct {
	name     string
	mana     int
	damage   int
	heal     int
	shield   int
	poison   int
	recharge int
}

var Spells = []Spell{
	{name: "Magic Missle", mana: 53, damage: 4},
	{name: "Drain", mana: 73, damage: 2, heal: 2},
	{name: "Shield", mana: 113, shield: 6},
	{name: "Poison", mana: 173, poison: 6},
	{name: "Recharge", mana: 229, recharge: 5},
}
