package day07

import (
	"adventofcode2015/pkg/assert"
	"log"
	"strconv"
	"strings"
)

type Wire struct {
	label       string
	value       int16
	initialized bool
}

type Wires map[string]*Wire

type Operation interface {
	// applies the operation to the wires, returns true if successful
	// or false if the we could not yet apply this operation
	Apply(Wires) bool
}

type Equal struct {
	input  string
	output string
}

func (op Equal) Apply(state Wires) bool {
	output := ensureWire(state, op.output)
	if output.initialized {
		return true
	}

	inputInt, initialized := inputValue(state, op.input)
	if !initialized {
		return false
	}

	output.value = inputInt
	output.initialized = true

	return true
}

type Not struct {
	input  string
	output string
}

func (op Not) Apply(state Wires) bool {
	output := ensureWire(state, op.output)
	if output.initialized {
		return true
	}

	inputInt, initialized := inputValue(state, op.input)
	if !initialized {
		return false
	}

	output.value = ^inputInt
	output.initialized = true
	return true
}

type Shift struct {
	input  string
	output string
	right  bool
	amount int16
}

func (op Shift) Apply(state Wires) bool {
	output := ensureWire(state, op.output)
	if output.initialized {
		return true
	}

	inputInt, initialized := inputValue(state, op.input)
	if !initialized {
		return false
	}

	switch op.right {
	case true:
		output.value = inputInt >> op.amount
	case false:
		output.value = inputInt << op.amount
	}
	output.initialized = true
	return true
}

type Calculate struct {
	left    string
	right   string
	command string
	output  string
}

func (op Calculate) Apply(state Wires) bool {
	output := ensureWire(state, op.output)

	if output.initialized {
		return true
	}

	leftInt, leftInitialized := inputValue(state, op.left)
	rightInt, rightInitialized := inputValue(state, op.right)

	if !leftInitialized || !rightInitialized {
		return false
	}

	switch op.command {
	case "AND":
		output.value = leftInt & rightInt
	case "OR":
		output.value = leftInt | rightInt
	default:
		log.Fatalf("unknown command: %s", op.command)
	}
	output.initialized = true
	return true
}

func NewOperation(line string) Operation {
	parts := strings.Split(line, " ")

	switch len(parts) {
	case 3:
		return Equal{parts[0], parts[2]}
	case 4:
		return Not{parts[1], parts[3]}
	default:
		switch parts[1] {
		case "LSHIFT":
			return Shift{parts[0], parts[4], false, mustAtoi(parts[2])}
		case "RSHIFT":
			return Shift{parts[0], parts[4], true, mustAtoi(parts[2])}
		default:
			return Calculate{parts[0], parts[2], parts[1], parts[4]}
		}
	}
}

func ensureWire(state Wires, label string) *Wire {
	existing, ok := state[label]
	if ok {
		return existing
	}
	wire := Wire{label, 0, false}
	state[label] = &wire
	return &wire
}

func inputValue(state Wires, label string) (int16, bool) {
	inputInt, err := strconv.ParseInt(label, 10, 16)

	// if we can't parse to a integer value then this must be a
	// directly assigned label
	if err != nil {
		inputWire := ensureWire(state, label)
		if !inputWire.initialized {
			return 0, false
		}
		return inputWire.value, true
	}

	return int16(inputInt), true
}

func mustAtoi(value string) int16 {
	out, err := strconv.ParseInt(value, 10, 16)
	assert.NoError(err, "could not convert value to int", value)
	return int16(out)
}
