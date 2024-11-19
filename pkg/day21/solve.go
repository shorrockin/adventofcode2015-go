package day21

import (
	"adventofcode2015/pkg/utils"
)

func PartOne(selfHitpoints, bossHitpoints, bossDamage, bossArmor int) int {
	for _, loadout := range LOADOUTS {
		if Simulate(selfHitpoints, loadout.Damage(), loadout.Armor(), bossHitpoints, bossDamage, bossArmor) {
			return loadout.Cost()
		}
	}
	return 0
}

func PartTwo(selfHitpoints, bossHitpoints, bossDamage, bossArmor int) int {
	for i := len(LOADOUTS) - 1; i >= 0; i-- {
		loadout := LOADOUTS[i]
		if !Simulate(selfHitpoints, loadout.Damage(), loadout.Armor(), bossHitpoints, bossDamage, bossArmor) {
			return loadout.Cost()
		}
	}
	return 0
}

func Simulate(selfHitpoints, selfDamage, selfArmor, bossHitpoints, bossDamage, bossArmor int) bool {
	for {
		damage := utils.MaxValue(1, selfDamage-bossArmor)
		bossHitpoints -= damage
		if bossHitpoints <= 0 {
			return true
		}

		damage = utils.MaxValue(1, bossDamage-selfArmor)
		selfHitpoints -= damage
		if selfHitpoints <= 0 {
			return false
		}
	}
}
