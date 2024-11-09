package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCalculation(t *testing.T) {
	op := NewOperation("hi AND hk -> hl").(Calculate)
	assert.Equal(t, "hi", op.left)
	assert.Equal(t, "hk", op.right)
	assert.Equal(t, "AND", op.command)
	assert.Equal(t, "hl", op.output)
}
