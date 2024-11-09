package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	result := Solve("part_one.example.txt")
	assert.Equal(t, int16(72), result["d"].value)
	assert.Equal(t, int16(507), result["e"].value)
	assert.Equal(t, int16(123), result["x"].value)
	assert.Equal(t, int16(456), result["y"].value)
}

func TestPartOne(t *testing.T) {
	result := Solve("part_one.input.txt")
	assert.Equal(t, int16(16076), result["a"].value)
}

func TestPartTwo(t *testing.T) {
	result := Solve("part_two.input.txt")
	assert.Equal(t, int16(2797), result["a"].value)
}
