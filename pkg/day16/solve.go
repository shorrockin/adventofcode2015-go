package day16

import (
	"adventofcode2015/pkg/utils"
	"strings"
)

var PROPERTIES map[string]int

func init() {
	PROPERTIES = map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

type Aunt struct {
	id        int
	compounds map[string]int
}

func (a Aunt) MatchesPartOne(properties map[string]int) bool {
	for name, value := range a.compounds {
		if properties[name] != value {
			return false
		}
	}
	return true
}

func (a Aunt) MatchesPartTwo(properties map[string]int) bool {
	for name, value := range a.compounds {
		switch name {
		case "cats", "trees":
			if properties[name] >= value {
				return false
			}
		case "pomeranians", "goldfish":
			if properties[name] <= value {
				return false
			}
		default:
			if properties[name] != value {
				return false
			}
		}
	}
	return true
}

func PartOne() int {
	return utils.Find(parse("input.txt"), func(aunt Aunt) bool {
		return aunt.MatchesPartOne(PROPERTIES)
	}).id
}

func PartTwo() int {
	return utils.Find(parse("input.txt"), func(aunt Aunt) bool {
		return aunt.MatchesPartTwo(PROPERTIES)
	}).id
}

func parse(path string) []Aunt {
	return utils.Map(utils.MustReadInput(path), func(line string) Aunt {
		parts := strings.Split(line, " ")
		compounds := make(map[string]int)
		for i := 2; i < len(parts); i += 2 {
			name := strings.TrimSuffix(parts[i], ":")
			value := strings.TrimSuffix(parts[i+1], ",")
			compounds[name] = utils.MustAtoi(value)
		}
		return Aunt{
			id:        utils.MustAtoi(strings.TrimSuffix(parts[1], ":")),
			compounds: compounds,
		}
	})
}
