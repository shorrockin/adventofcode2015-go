package day11

import "testing"
import "github.com/stretchr/testify/assert"

func Next(value string) string {
	return string(NextRune([]rune(value)))
}

func TestNextRune(t *testing.T) {
	assert.Equal(t, "b", Next("a"))
	assert.Equal(t, "j", Next("h"))
	assert.Equal(t, "ac", Next("ab"))
	assert.Equal(t, "ba", Next("az"))
	assert.Equal(t, "aa", Next("z"))
	assert.Equal(t, "ya", Next("xz"))
}

func TestValidatePassword(t *testing.T) {
	assert.True(t, ValidPassword([]rune("abcddee")))
	assert.True(t, ValidPassword([]rune("zabctffaa")))
	assert.False(t, ValidPassword([]rune("abbceffg")))
	assert.False(t, ValidPassword([]rune("abbcegjk")))
	assert.False(t, ValidPassword([]rune("abbcegjkbb")))
	assert.True(t, ValidPassword([]rune("abbcdgjkcc")))
}

func TestNextPasswordExample(t *testing.T) {
	assert.Equal(t, "abcdffaa", NextPassword("abcdefgh"))
}

func TestNextPasswordActual(t *testing.T) {
	assert.Equal(t, "cqjxxyzz", NextPassword("cqjxjnds"))
	assert.Equal(t, "cqkaabcc", NextPassword("cqjxxyzz"))
}
