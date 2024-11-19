package day20

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExampleInput(t *testing.T) {
	assert.Equal(t, 8, Solve(130, 10, math.MaxInt))
}

func TestPartOneActualInput(t *testing.T) {
	assert.Equal(t, 776160, Solve(33100000, 10, math.MaxInt))
}

func TestPartTwoActualInput(t *testing.T) {
	assert.Equal(t, 786240, Solve(33100000, 11, 50))
}
