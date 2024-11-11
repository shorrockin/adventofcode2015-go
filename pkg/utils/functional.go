package utils

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
