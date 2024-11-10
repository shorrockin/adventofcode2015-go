package day09

import (
	"adventofcode2015/pkg/graph"
	"adventofcode2015/pkg/utils"
	"strings"
)

func Solve(path string, shortest bool) float64 {
	lines := utils.MustReadInput(path)
	data := graph.NewGraph[string]()

	for _, line := range lines {
		parts := strings.Split(line, " ")
		distance := utils.MustAtoi(parts[4])
		data.AddBidirectionalEdge(parts[0], parts[2], float64(distance))
	}

	solution := graph.TSP(data, shortest)
	return solution.Distance()
}
