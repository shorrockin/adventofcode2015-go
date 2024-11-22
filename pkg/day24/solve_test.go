package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExampleInput(t *testing.T) {
	assert.Equal(t, 99, Solve("input.example.txt", 3))
}

func TestPartOneActualInput(t *testing.T) {
	assert.Equal(t, 10439961859, Solve("input.txt", 3))
}

func TestPartTwoExampleInput(t *testing.T) {
	assert.Equal(t, 44, Solve("input.example.txt", 4))
}

func TestPartTwoActualInput(t *testing.T) {
	assert.Equal(t, 72050269, Solve("input.txt", 4))
}
