package day13

import "testing"
import "github.com/stretchr/testify/assert"

func TestSolvePartOneExample(t *testing.T) {
	assert.Equal(t, 330, Solve("example.txt", false))
}

func TestSolvePartOneActual(t *testing.T) {
	assert.Equal(t, 709, Solve("input.txt", false))
}

func TestSolvePartTwoActual(t *testing.T) {
	assert.Equal(t, 668, Solve("input.txt", true))
}
