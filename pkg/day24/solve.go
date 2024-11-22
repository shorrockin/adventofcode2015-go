package day24

import (
	"adventofcode2015/pkg/utils"
)

func Solve(path string, buckets int) int {
	data := parse(path)
	weight := sum(data) / buckets

	for i := 1; i < len(data); i++ {
		for _, bucket := range utils.Combinations(data, i) {
			var best []int

			if sum(bucket) == weight {
				if len(best) == 0 || product(best) < product(bucket) {
					best = bucket
				}
			}

			if len(best) > 0 {
				return product(best)
			}
		}
	}

	return 0
}

func parse(path string) []int {
	return utils.Map(utils.MustReadInput(path), func(line string) int {
		return utils.MustAtoi(line)
	})
}

func sum(values []int) int {
	return utils.Reduce(values, 0, func(accum int, next int) int {
		return accum + next
	})
}

func product(values []int) int {
	return utils.Reduce(values, 1, func(accum int, next int) int {
		return accum * next
	})
}
