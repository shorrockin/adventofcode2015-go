package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneSimulation(t *testing.T) {
	assert.True(t, Simulate(8, 5, 5, 12, 7, 2))
}

func TestPartOneActualInput(t *testing.T) {
	assert.Equal(t, 78, PartOne(100, 104, 8, 1))
}

func TestPartTwoActualInput(t *testing.T) {
	assert.Equal(t, 148, PartTwo(100, 104, 8, 1))
}
