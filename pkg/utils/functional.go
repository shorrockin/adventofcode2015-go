package utils

import "golang.org/x/exp/constraints"

func Map[T any, K any](data []T, mapper func(T) K) []K {
	result := make([]K, len(data))
	for index, element := range data {
		result[index] = mapper(element)
	}
	return result
}

func Uniq[T comparable](input []T) []T {
	seen := NewSet[T]()
	var result []T

	for _, value := range input {
		if !seen.Contains(value) {
			result = append(result, value)
			seen.Add(value)
		}
	}

	return result
}

func Max[T constraints.Ordered](input []T) T {
	if len(input) == 0 {
		var zero T
		return zero
	}

	best := input[0]
	for _, current := range input {
		if current > best {
			best = current
		}
	}

	return best
}

func MaxMapValue[K comparable, V constraints.Ordered](input map[K]V) V {
	var zero V
	if len(input) == 0 {
		return zero
	}

	best := zero
	for _, value := range input {
		if value > best {
			best = value
		}
	}

	return best
}
