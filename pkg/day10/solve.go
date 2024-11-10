package day10

import (
	"adventofcode2015/pkg/assert"
	"strconv"
	"strings"
)

func Solve(input string, iterations int) string {
	assert.Assert(len(input) > 0, "must have a non-zero length string")

	var builder strings.Builder
	index := 1
	count := 1
	character := input[0]

	for index < len(input) {
		// if the next character is same as the last, then we just
		// add that to our character count and move on to the next
		// value.
		if input[index] == character {
			count++
			index++
			continue
		}

		builder.WriteString(strconv.Itoa(count))
		builder.WriteString(string(character))

		count = 1
		character = input[index]

		index++
	}

	builder.WriteString(strconv.Itoa(count))
	builder.WriteString(string(character))

	if iterations > 1 {
		return Solve(builder.String(), iterations-1)
	}

	return builder.String()
}
