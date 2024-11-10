package day10

import "testing"
import "github.com/stretchr/testify/assert"

func TestSolvePartOneExampleInput(t *testing.T) {
	assert.Equal(t, "11", Solve("1", 1))
	assert.Equal(t, "21", Solve("11", 1))
	assert.Equal(t, "1211", Solve("21", 1))
	assert.Equal(t, "111221", Solve("1211", 1))
	assert.Equal(t, "312211", Solve("111221", 1))
	assert.Equal(t, "312211", Solve("1", 5))
}

func TestSolvePartOneActual(t *testing.T) {
	assert.Equal(t, int(360154), len(Solve("1113122113", 40)))
}

func TestSolvePartTwoActual(t *testing.T) {
	assert.Equal(t, int(5103798), len(Solve("1113122113", 50)))
}
