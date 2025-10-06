package gocalc

import (
	"math"
)

var constants = map[string]float64{
	"pi":  math.Pi,
	"phi": math.Phi,
	"e":   math.E,
}

type Program []Instruction

type Instruction func(stack *Stack) error

func constant(constName string) Instruction {
	return func(stack *Stack) error {
		constValue := constants[constName]
		return stack.Push(StackData(constValue))
	}
}

func literal(value StackData) Instruction {
	return func(stack *Stack) error {
		return stack.Push(value)
	}
}

func add(stack *Stack) error {
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
	result := op1 + op2
	return stack.Push(result)
}

func sub(stack *Stack) error {
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
	result := op1 - op2
	return stack.Push(result)
}

func mul(stack *Stack) error {
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
	result := op1 * op2
	return stack.Push(result)
}

func div(stack *Stack) error {
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
	result := op1 / op2
	return stack.Push(result)
}
