package day07

import (
	"adventofcode2015/pkg/utils"
)

func Solve(path string) Wires {
	lines := utils.MustReadInput(path)

	var operations []Operation

	for _, str := range lines {
		operations = append(operations, NewOperation(str))
	}

	var wires Wires = make(Wires)
	var complete = false

	for !complete {
		complete = true
		for _, op := range operations {
			local := op.Apply(wires)
			complete = local && complete
		}
	}

	return wires
}
