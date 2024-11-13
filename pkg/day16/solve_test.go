package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, 40, PartOne())
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 241, PartTwo())
}
