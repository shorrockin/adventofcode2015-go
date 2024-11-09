package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 12, PartOne("example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1333, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 19, PartTwo("example.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 2046, PartTwo("input.txt"))
}
