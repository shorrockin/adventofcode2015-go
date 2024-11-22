package utils

func Combinations[T any](data []T, size int) [][]T {
	var result [][]T

	var helper func(start int, current []T)
	helper = func(start int, current []T) {
		if len(current) == size {
			tmp := make([]T, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		for i := start; i < len(data); i++ {
			helper(i+1, append(current, data[i]))
		}
	}

	helper(0, []T{})
	return result
}
