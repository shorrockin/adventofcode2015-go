package day19

import "testing"
import "github.com/stretchr/testify/assert"

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 7, PartOne("example.partone.txt"))
}

func TestPartOneReplacementVariations(t *testing.T) {
	assert.Equal(t, "HOOH", replace("HOH", 0, 1, "HO"))
	assert.Equal(t, "HOHO", replace("HOH", 2, 1, "HO"))
	assert.Equal(t, "HHHH", replace("HOH", 1, 1, "HH"))
	assert.Equal(t, "HHH", replace("HOH", 1, 2, "HH"))
	assert.Equal(t, "OHH", replace("HOH", 0, 2, "OH"))
	assert.Equal(t, "HABCD", replace("HOH", 1, 2, "ABCD"))
	assert.Equal(t, "HABCDEFG", replace("HOHEFG", 1, 2, "ABCD"))
	assert.Equal(t, "ABCDHEFG", replace("HOHEFG", 0, 2, "ABCD"))
	assert.Equal(t, "HOG", replace("HOHEFG", 1, 4, "O"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 535, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 6, PartTwo("example.parttwo.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 212, PartTwo("input.txt"))
}
