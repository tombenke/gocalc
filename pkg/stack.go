package gocalc

import (
	"fmt"
)

type StackData float64

type Stack struct {
	data    []StackData
	pointer int
	size    int
}

func NewStack(size int) Stack {
	stack := Stack{data: make([]StackData, size), size: size, pointer: -1}
	return stack
}

func (s *Stack) Push(data StackData) error {
	if s.pointer >= s.size {
		return fmt.Errorf("stack overflow")
	}
	s.pointer++
	s.data[s.pointer] = data
	return nil
}

func (s *Stack) Pop() (StackData, error) {
	if s.pointer < 0 {
		return 0., fmt.Errorf("stack is empty")
	}
	rval := s.data[s.pointer]
	s.pointer--
	return rval, nil
}

func (s *Stack) GetPointer() int {
	return s.pointer
}
