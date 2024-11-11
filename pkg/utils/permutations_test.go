package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	input := []int{1, 2, 3}
	permutations := Permutations(input)
	assert.Equal(t, 6, len(permutations))
	assert.Equal(t, []int{1, 2, 3}, permutations[0])
	assert.Equal(t, []int{3, 1, 2}, permutations[5])
}
