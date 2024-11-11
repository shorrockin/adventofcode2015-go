package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 1120, PartOne("part_one.example.txt", 1000))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 2660, PartOne("part_one.input.txt", 2503))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 689, PartTwo("part_one.example.txt", 1000))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1256, PartTwo("part_one.input.txt", 2503))
}
