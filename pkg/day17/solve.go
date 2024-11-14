package day17

import (
	"adventofcode2015/pkg/utils"
	"sort"
)

func Solve(path string, capacity int, partOne bool) int {
	if partOne {
		return countContainers(parse(path), 0, capacity, 0, 0)
	}
	containers, _ := countMinContainers(parse(path), 0, capacity, 0, 99999, 1, 0)
	return containers
}

func countContainers(containers []int, start int, maxCapacity int, capacityUsed int, result int) int {
	for index := start; index < len(containers); index++ {
		containerSize := containers[index]
		if capacityUsed+containerSize > maxCapacity {
			return result
		}
		if capacityUsed+containerSize == maxCapacity {
			result = result + 1
		}
		result = countContainers(containers, index+1, maxCapacity, capacityUsed+containerSize, result)
	}
	return result
}

func countMinContainers(containers []int, start int, maxCapacity int, capacityUsed int, minContainersUsed int, containersUsed int, result int) (int, int) {
	// we've used more containers than our min containers we
	// can short circuit
	if containersUsed > minContainersUsed {
		return result, minContainersUsed
	}

	for index := start; index < len(containers); index++ {
		containerSize := containers[index]
		if capacityUsed+containerSize > maxCapacity {
			return result, minContainersUsed
		}
		if capacityUsed+containerSize == maxCapacity {
			if containersUsed < minContainersUsed {
				minContainersUsed = containersUsed
				result = 1
			} else {
				result = result + 1
			}
		}

		result, minContainersUsed = countMinContainers(containers, index+1, maxCapacity, capacityUsed+containerSize, minContainersUsed, containersUsed+1, result)
	}
	return result, minContainersUsed
}

func parse(path string) []int {
	values := utils.Map(utils.MustReadInput(path), func(line string) int {
		return utils.MustAtoi(line)
	})
	sort.Ints(values)
	return values
}
