package day13

import (
	"adventofcode2015/pkg/assert"
	"adventofcode2015/pkg/utils"
	"strconv"
	"strings"
)

type Rule struct {
	name     string
	amount   int
	neighbor string
}

type Rules map[string]map[string]Rule

func Solve(path string, includeSelf bool) int {
	rulesSlice := parse(path)
	rules := toRules(rulesSlice)

	names := utils.Map(rulesSlice, func(rule Rule) string { return rule.name })
	names = utils.Uniq(names)
	if includeSelf {
		names = append(names, "Self")
	}

	best := 0
	for _, arrangement := range utils.Permutations(names) {
		score := score(arrangement, rules)
		if score > best {
			best = score
		}
	}

	return best
}

func parse(path string) []Rule {
	return utils.Map(utils.MustReadInput(path), func(line string) Rule {
		parts := strings.Split(line, " ")
		amount, err := strconv.Atoi(parts[3])
		assert.NoError(err, "could not convert amount to int", parts[3])

		if parts[2] == "lose" {
			amount = amount * -1
		}

		return Rule{
			name:     parts[0],
			amount:   amount,
			neighbor: parts[10][:len(parts[10])-1],
		}
	})
}

func toRules(rules []Rule) Rules {
	out := make(Rules)
	for _, rule := range rules {
		_, exists := out[rule.name]
		if !exists {
			out[rule.name] = make(map[string]Rule)
		}
		out[rule.name][rule.neighbor] = rule
	}
	return out
}

func score(arrangment []string, rules Rules) int {
	score := 0
	for index, name := range arrangment {
		// for people to the right of us:
		if index == len(arrangment)-1 {
			score += rules[name][arrangment[0]].amount
		} else {
			score += rules[name][arrangment[index+1]].amount
		}

		// for people to the left of us:
		if index == 0 {
			score += rules[name][arrangment[len(arrangment)-1]].amount
		} else {
			score += rules[name][arrangment[index-1]].amount
		}
	}

	return score
}
