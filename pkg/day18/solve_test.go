package day18

import "testing"
import "github.com/stretchr/testify/assert"

func TestPartOneExampleInput(t *testing.T) {
	assert.Equal(t, 4, Solve("example.txt", 4, false))
}

func TestPartOneActualInput(t *testing.T) {
	assert.Equal(t, 814, Solve("input.txt", 100, false))
}

func TestPartTwoExampleInput(t *testing.T) {
	assert.Equal(t, 17, Solve("example.txt", 5, true))
}

func TestPartTwoActualInput(t *testing.T) {
	assert.Equal(t, 924, Solve("input.txt", 100, true))
}
