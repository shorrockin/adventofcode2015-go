package day18

import (
	"adventofcode2015/pkg/grid"
	"adventofcode2015/pkg/utils"
)

// simple map containing a time mapped to
// if the light is on or off at that specific
// moment in time
type History map[int]bool

func Solve(path string, iterations int, partTwo bool) int {
	lights := parse(path)
	stuck := utils.NewSet[grid.Coordinate]()

	if partTwo {
		_, width := lights.Width()
		_, height := lights.Height()
		stuck.Add(grid.Coordinate{X: 0, Y: 0})
		stuck.Add(grid.Coordinate{X: width, Y: 0})
		stuck.Add(grid.Coordinate{X: 0, Y: height})
		stuck.Add(grid.Coordinate{X: width, Y: height})
		for coordinate := range stuck {
			lights.MustGet(coordinate).Contents[0] = true
		}
	}

	for i := 0; i <= iterations; i++ {
		for coordinate, value := range lights {
			neighbors := lights.GetAll(coordinate.Compass())
			neighbors = utils.Filter(neighbors, func(node grid.Node[History]) bool {
				return node.Contents[i]
			})
			lightsOn := len(neighbors)

			if stuck.Contains(coordinate) {
				value.Contents[i+1] = true
			} else if value.Contents[i] {
				value.Contents[i+1] = (lightsOn == 2 || lightsOn == 3)
			} else {
				value.Contents[i+1] = (lightsOn == 3)
			}
		}
	}

	return len(utils.FilterMap(lights, func(coord grid.Coordinate, node grid.Node[History]) bool {
		return node.Contents[iterations]
	}))
}

func toRune(history History) rune {
	size := len(history)
	if history[size-1] {
		return '#'
	}
	return '.'
}

func parse(path string) grid.Grid[History] {
	return grid.Parse(utils.MustReadInput(path), func(value rune, x int, y int) History {
		element := make(map[int]bool)
		element[0] = value == '#'
		return element
	})
}
