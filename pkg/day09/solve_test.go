package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOneExample(t *testing.T) {
	distance := Solve("example.txt", true)
	assert.Equal(t, float64(605), distance)
}

func TestSolvePartOneActual(t *testing.T) {
	distance := Solve("input.txt", true)
	assert.Equal(t, float64(207), distance)
}

func TestSolvePartTwo(t *testing.T) {
	distance := Solve("example.txt", false)
	assert.Equal(t, float64(982), distance)
}

func TestSolvePartTwoActual(t *testing.T) {
	distance := Solve("input.txt", false)
	assert.Equal(t, float64(804), distance)
}
