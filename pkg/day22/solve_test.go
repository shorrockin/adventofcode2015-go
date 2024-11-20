package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 900, Solve(false))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 1216, Solve(true))
}
