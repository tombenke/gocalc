package gocalc

import (
	"fmt"
	"github.com/tombenke/parc"
	"sync"
)

const (
	// The initial size of the data stack
	DATA_STACK_INIT_SIZE = 100
)

// GoCalc is a struct that holds a data stack, a program memory, and an instruction pointer.
// It implements a minimalistic stack machine that runs the instructions stored in the program memory.
// The instructions are function pointers. They use the data stack to manipulate data,
// and to communicate with each other.
type GoCalc struct {
	// mu prevents the conflict between compile and run activities of the calculator
	mu *sync.Mutex

	// dataStack holds the stack for data items manipulated by the instructions of the program
	dataStack Stack

	// program stores the instructions of the program of the calculator
	program Program

	// debug holds the complete program in string format for debugging purposes
	debug []string

	// Instruction Pointer. It is an index on the program, and points to the next instruction to execute
	ip int

	// ast is a kind of Abstract Syntax Tree. It is the results of the last parsing.
	ast parc.Result
}

// NewGoCalc Create a new instance of GoCalc
func NewGoCalc() (*GoCalc, error) {
	return &GoCalc{dataStack: NewStack(DATA_STACK_INIT_SIZE), program: make([]Instruction, 100), debug: []string{}, ip: 0, mu: &sync.Mutex{}}, nil
}

// GetIP returns with the actual value of the Instruction Pointer
func (c GoCalc) GetIP() int {
	return c.ip
}

// GetIP returns with the actual value of the Stack Pointer of the Data Stack
func (c GoCalc) GetDataStackPointer() int {
	return c.dataStack.GetPointer()
}

func (c GoCalc) GetAST() parc.Result {
	return c.ast
}

func (c GoCalc) GetProgramDebug() []string {
	return c.debug
}

// Compile parses the `source` code, and generates the corresponding sequence of instructions into the `program` store
func (c *GoCalc) Compile(source string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	// Make sure to have an end of an empty program
	c.program[0] = nil
	c.debug = c.debug[:0]

	// Parse the source code
	parseResultAST, parseError := Parse(source)
	if parseError != nil {
		panic("parsing error")
	}

	c.ast = parseResultAST

	// Encode the parsed results into a series of instructions
	nextInstruction := c.encode(c.ast, 0)
	c.program[nextInstruction] = nil
	c.debug = append(c.debug, "NIL")
}

// encode converts a node from the parsing results into an intruction, that also places ito the program.
func (c *GoCalc) encode(node parc.Result, nextInstruction int) int {
	switch n := node.(type) {
	case Term:
		c.debug = append(c.debug, "# Term")
		nextInstruction = c.encode(n.Operand_A, nextInstruction)
		nextInstruction = c.encode(n.Operand_B, nextInstruction)
		nextInstruction = c.encode(n.Operator, nextInstruction)

	case Expression:
		c.debug = append(c.debug, "# Expression")
		nextInstruction = c.encode(n.Operand_A, nextInstruction)
		nextInstruction = c.encode(n.Operand_B, nextInstruction)
		nextInstruction = c.encode(n.Operator, nextInstruction)

	case Operator:
		switch n.Value {
		case "+":
			c.program[nextInstruction] = add
			c.debug = append(c.debug, "ADD")
		case "-":
			c.program[nextInstruction] = sub
			c.debug = append(c.debug, "SUB")
		case "/":
			c.program[nextInstruction] = div
			c.debug = append(c.debug, "DIV")
		case "*":
			c.program[nextInstruction] = mul
			c.debug = append(c.debug, "MUL")
		}
		nextInstruction++

	case Number:
		c.program[nextInstruction] = literal(&(c.dataStack), n.Value)
		c.debug = append(c.debug, fmt.Sprintf("LIT %v\n", n.Value))
		nextInstruction++
	}
	return nextInstruction
}

// Run executes the instructions stored in the program memory.
// The execution begins with the 0th instruction,
// and the Instruction Pointer points to the next instruction to be executed.
// The result of the program will be the top item on the data stack, if there is any.
// If the stack is empty then the result will be nil
func (c GoCalc) Run() *StackData {
	c.mu.Lock()
	defer c.mu.Unlock()

	for {
		if c.program[c.ip] == nil {
			break
		}
		err := c.program[c.ip](&(c.dataStack))
		if err != nil {
			panic(fmt.Sprintf("GoCalc internal error in executing at IP: %d", c.ip))
		}
		c.ip++
	}

	// Check if Data Stack is empty
	if c.GetDataStackPointer() < 0 {
		return nil
	}

	result, err := c.dataStack.Pop()
	if err != nil {
		panic(fmt.Sprintf("GoCalc internal error in getting the result of the program IP: %d", c.ip))
	}

	return &result
}
