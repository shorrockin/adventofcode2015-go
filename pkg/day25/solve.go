package day25

func Solve(rowTarget, columnTarget int) int {
	value := 20151125
	row := 1
	column := 1

	for {
		row -= 1
		column += 1

		if row == 0 {
			row = column
			column = 1
		}

		value = (value * 252533) % 33554393
		if column == columnTarget && row == rowTarget {
			return value
		}
	}
}
