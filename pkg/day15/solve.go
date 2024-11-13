package day15

import (
	"adventofcode2015/pkg/assert"
	"adventofcode2015/pkg/utils"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Ingredients []Ingredient

type Amounts map[Ingredient]int

func (amounts *Amounts) Score() int {
	capacity := amounts.reduce(func(i Ingredient) int { return i.capacity })
	durability := amounts.reduce(func(i Ingredient) int { return i.durability })
	flavor := amounts.reduce(func(i Ingredient) int { return i.flavor })
	texture := amounts.reduce(func(i Ingredient) int { return i.texture })
	return capacity * durability * flavor * texture
}

func (amounts *Amounts) Calories() int {
	return amounts.reduce(func(i Ingredient) int { return i.calories })
}

func (amounts *Amounts) reduce(getter func(i Ingredient) int) int {
	accum := 0
	for ingredient, amount := range *amounts {
		accum += getter(ingredient) * amount
	}
	return utils.MaxValue(0, accum)
}

func PartOne(path string) int {
	amounts := make(Amounts)
	return solve(parse(path), &amounts, 0, 100, 0, true)
}

func PartTwo(path string) int {
	amounts := make(Amounts)
	return solve(parse(path), &amounts, 0, 100, 0, false)
}

func solve(ingredients Ingredients, amounts *Amounts, index int, remaining int, best int, partOne bool) int {
	ingredient := ingredients[index]

	// we're at the last ingredient, we need to figure out
	// what the score is.
	if index == len(ingredients)-1 {
		(*amounts)[ingredient] = remaining

		if !partOne {
			calories := amounts.Calories()
			if calories != 500 {
				return best
			}
		}

		return utils.MaxValue(amounts.Score(), best)
	}

	for i := 0; i < remaining; i++ {
		(*amounts)[ingredient] = i
		best = solve(ingredients, amounts, index+1, remaining-i, best, partOne)
	}

	return best
}

func parse(path string) Ingredients {
	chop := func(value string) string { return value[:len(value)-1] }
	return utils.Map(utils.MustReadInput(path), func(line string) Ingredient {
		parts := strings.Split(line, " ")
		assert.Assert(11 == len(parts), "expected 11 parts")
		return Ingredient{
			name:       chop(parts[0]),
			capacity:   utils.MustAtoi(chop(parts[2])),
			durability: utils.MustAtoi(chop(parts[4])),
			flavor:     utils.MustAtoi(chop(parts[6])),
			texture:    utils.MustAtoi(chop(parts[8])),
			calories:   utils.MustAtoi(parts[10]),
		}
	})

}
