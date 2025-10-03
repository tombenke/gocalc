package gocalc

import (
	"fmt"
	"math"
)

var constants = map[string]float64{
	"pi":  math.Pi,
	"phi": math.Phi,
	"e":   math.E,
}

type Program []Instruction

type Instruction func(stack *Stack) error

func constant(stack *Stack, constName string) Instruction {
	return func(stack *Stack) error {
		constValue := constants[constName]
		fmt.Printf("constant %s => %.2f\n", constName, constValue)
		return stack.Push(StackData(constValue))
	}
}

func literal(stack *Stack, value StackData) Instruction {
	return func(stack *Stack) error {
		//fmt.Printf("literal called with %+v\n", value)
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
