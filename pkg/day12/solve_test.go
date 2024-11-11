package day12

import (
	"adventofcode2015/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneExample(t *testing.T) {
	mustSolve := func(json string) float64 {
		result, err := Solve(json, false)
		assert.NoError(t, err)
		return result
	}

	assert.Equal(t, float64(6), mustSolve(`{"a":2,"b":4}`))
	assert.Equal(t, float64(6), mustSolve(`[1,2,3]`))
	assert.Equal(t, float64(0), mustSolve(`[-1,{"a":1}]`))
}

func TestSolvePartOneActual(t *testing.T) {
	json := utils.MustReadInput("part_one.input.txt")
	result, err := Solve(json[0], false)
	assert.NoError(t, err)
	assert.Equal(t, float64(111754), result)
}

func TestSolvePartTwoActual(t *testing.T) {
	json := utils.MustReadInput("part_one.input.txt")
	result, err := Solve(json[0], true)
	assert.NoError(t, err)
	assert.Equal(t, float64(65402), result)
}
