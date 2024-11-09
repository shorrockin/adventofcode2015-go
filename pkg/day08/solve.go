package day08

import (
	"adventofcode2015/pkg/assert"
	"adventofcode2015/pkg/utils"
	"strconv"
)

func PartOne(path string) int {
	lines := utils.MustReadInput(path)

	var count = 0
	for _, line := range lines {
		escaped, err := strconv.Unquote(line)
		assert.NoError(err, "unable to unquote string", line)
		count += len(line) - len(escaped)
	}

	return count
}

func PartTwo(path string) int {
	lines := utils.MustReadInput(path)
	var count = 0
	for _, line := range lines {
		encoded := strconv.Quote(line)
		count += len(encoded) - len(line)
	}

	return count
}
