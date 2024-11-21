package day23

import (
	"adventofcode2015/pkg/utils"
	"log"
	"strings"
)

func Solve(path string, a int) Registers {
	registers := Registers{a: a}
	operations := parse(path)
	next := 0

	for {
		next += operations[next](&registers)

		if len(operations) <= next {
			return registers
		}
	}
}

func parse(path string) []Operation {
	lines := utils.MustReadInput(path)
	operations := make([]Operation, 0, len(lines))

	toRegisterId := func(value string) RegisterID {
		switch value[0] {
		case 'a':
			return A
		case 'b':
			return B
		default:
			log.Fatalf("unable to get the register id of value: %v", value)
			return -1
		}
	}

	toOffset := func(value string) int {
		switch value[0] {
		case '+':
			return utils.MustAtoi(value[1:])
		case '-':
			return utils.MustAtoi(value[1:]) * -1
		default:
			log.Fatalf("expected offset to start with -/+: %v", value)
			return 0
		}
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "hlf":
			operations = append(operations, Half(toRegisterId(parts[1])))
		case "tpl":
			operations = append(operations, Triple(toRegisterId(parts[1])))
		case "inc":
			operations = append(operations, Increment(toRegisterId(parts[1])))
		case "jmp":
			operations = append(operations, Jump(toOffset(parts[1])))
		case "jie":
			operations = append(operations, JumpIfEven(toRegisterId(parts[1]), toOffset(parts[2])))
		case "jio":
			operations = append(operations, JumpIfOne(toRegisterId(parts[1]), toOffset(parts[2])))
		default:
			log.Fatalf("unable to parse: %v", line)
		}
	}

	return operations
}
