package utils

import (
	"adventofcode2015/pkg/assert"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func MustReadInput(path string) []string {
	lines, err := ReadInput(path)
	if err != nil {
		log.Fatalf("Could not ReadInput: '%v': err: %+v", path, err)
	}
	return lines
}

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func MustAtoi(raw string) int {
	value, err := strconv.Atoi(raw)
	assert.NoError(err, "could not convert value to number", raw)
	return value
}

func Must[T any](value T, err error) T {
	assert.NoError(err, "Must passed value with err: %v", err)
	return value
}

func CopySlice[T any](original []T) []T {
	copied := make([]T, len(original))
	copy(copied, original)
	return copied
}

func CopyMap[K comparable, V any](original map[K]V) map[K]V {
	copy := make(map[K]V, len(original))
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func Indexes(value string, target string) []int {
	indexes := []int{}
	offset := 0

	for {
		index := strings.Index(value[offset:], target)
		if index == -1 {
			break
		}

		indexes = append(indexes, offset+index)
		offset = offset + index + 1
	}

	return indexes
}
