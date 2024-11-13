package day15

import "testing"
import "github.com/stretchr/testify/assert"

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 62842880, PartOne("example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 18965440, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 57600000, PartTwo("example.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 15862900, PartTwo("input.txt"))
}
