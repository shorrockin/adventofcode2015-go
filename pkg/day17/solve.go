package day17

import (
	"adventofcode2015/pkg/utils"
	"sort"
)

func Solve(path string, capacity int) int {
	return solve(parse(path), 0, capacity, 0, 0)
}

func solve(containers []int, start int, capacity int, used int, result int) int {
	for index := start; index < len(containers); index++ {
		container := containers[index]

		// no need to keep iterating
		if used+container > capacity {
			return result
		}

		// did we hit capacity? add to result, and keep the loop
		// going as we may hit a duplicate in following iterations
		if used+container == capacity {
			result = result + 1
		}
		result = solve(containers, index+1, capacity, used+container, result)
	}
	return result
}

func parse(path string) []int {
	values := utils.Map(utils.MustReadInput(path), func(line string) int {
		return utils.MustAtoi(line)
	})
	sort.Ints(values)
	return values
}
