package day23

import "adventofcode2015/pkg/assert"

type RegisterID int

const (
	A RegisterID = iota
	B
)

func AssertValidRegisterID(id RegisterID) {
	assert.True(id == A || id == B, "register only supports A & B", id)
}

type Registers struct {
	a int
	b int
}

func (r *Registers) Get(id RegisterID) int {
	AssertValidRegisterID(id)
	switch id {
	case A:
		return r.a
	default:
		return r.b
	}
}

func (r *Registers) Set(id RegisterID, value int) {
	assert.True(value >= 0, "register value can not be negative", value)
	AssertValidRegisterID(id)
	switch id {
	case A:
		r.a = value
	default:
		r.b = value
	}
}

type Operation func(*Registers) int

func Half(id RegisterID) Operation {
	return func(registers *Registers) int {
		registers.Set(id, registers.Get(id)/2)
		return 1
	}
}

func Triple(id RegisterID) Operation {
	return func(registers *Registers) int {
		registers.Set(id, registers.Get(id)*3)
		return 1
	}
}

func Increment(id RegisterID) Operation {
	return func(registers *Registers) int {
		registers.Set(id, registers.Get(id)+1)
		return 1
	}
}

func Jump(offset int) Operation {
	return func(registers *Registers) int {
		return offset
	}
}

func JumpIfEven(id RegisterID, offset int) Operation {
	return func(registers *Registers) int {
		if registers.Get(id)%2 == 0 {
			return offset
		}
		return 1
	}
}

func JumpIfOne(id RegisterID, offset int) Operation {
	return func(registers *Registers) int {
		if registers.Get(id) == 1 {
			return offset
		}
		return 1
	}
}
