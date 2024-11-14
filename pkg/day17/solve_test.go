package day17

import "testing"
import "github.com/stretchr/testify/assert"

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 4, Solve("input.example.txt", 25, true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1304, Solve("input.txt", 150, true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 3, Solve("input.example.txt", 25, false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 18, Solve("input.txt", 150, false))
}
