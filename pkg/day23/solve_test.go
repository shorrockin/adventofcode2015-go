package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneExampleInput(t *testing.T) {
	assert.Equal(t, 2, Solve("input.example.txt", 0).a)
}

func TestSolvePartOneInput(t *testing.T) {
	assert.Equal(t, 255, Solve("input.txt", 0).b)
}

func TestSolvePartTwoInput(t *testing.T) {
	assert.Equal(t, 334, Solve("input.txt", 1).b)
}
