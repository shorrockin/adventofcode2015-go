package day21

import (
	"fmt"
	"sort"
)

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type Loadout struct {
	Weapon    Item
	Chest     Item
	RingRight Item
	RingLeft  Item
}

func (l Loadout) Cost() int {
	return l.Weapon.cost + l.Chest.cost + l.RingRight.cost + l.RingLeft.cost
}
func (l Loadout) Damage() int {
	return l.Weapon.damage + l.Chest.damage + l.RingRight.damage + l.RingLeft.damage
}
func (l Loadout) Armor() int {
	return l.Weapon.armor + l.Chest.armor + l.RingRight.armor + l.RingLeft.armor
}
func (l Loadout) String() string {
	return fmt.Sprintf(
		"%v + %v + %v + %v (cost=%v, damage=%v, armor=%v)",
		l.Weapon.name,
		l.Chest.name,
		l.RingRight.name,
		l.RingLeft.name,
		l.Cost(),
		l.Damage(),
		l.Armor(),
	)

}

var WEAPONS []Item
var ARMOR []Item
var RINGS []Item
var LOADOUTS []Loadout

func init() {
	WEAPONS = []Item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	ARMOR = []Item{
		{"None", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}

	RINGS = []Item{
		{"None", 0, 0, 0},
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}

	for _, weapon := range WEAPONS {
		for _, armor := range ARMOR {
			for _, rightRing := range RINGS {
				for ringIdx, leftRing := range RINGS {
					if rightRing == leftRing && ringIdx != 0 {
						continue
					}

					LOADOUTS = append(LOADOUTS, Loadout{
						Weapon:    weapon,
						Chest:     armor,
						RingRight: rightRing,
						RingLeft:  leftRing,
					})
				}
			}
		}
	}

	sort.Slice(LOADOUTS, func(left, right int) bool {
		return LOADOUTS[left].Cost() < LOADOUTS[right].Cost()
	})
}
