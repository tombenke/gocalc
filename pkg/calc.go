package gocalc

import (
	"fmt"
	"github.com/tombenke/gocalc/pkg/buildinfo"
	"github.com/tombenke/parc"
	"io"
	"os"
	"sync"
)

const (
	// The initial size of the data stack
	DATA_STACK_INIT_SIZE = 100
)

var debugWriter io.Writer = io.Discard

func init() {
	fmt.Printf("BuildTags: %+v\nDebugMode: %v\n", buildinfo.BuildTags, buildinfo.DebugMode)
	if buildinfo.DebugMode {
		debugWriter = os.Stdout
	}
}

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

	// Instruction Pointer. It is an index on the program, and points to the next instruction to execute
	ip int
}

// NewGoCalc Create a new instance of GoCalc
func NewGoCalc() (*GoCalc, error) {
	return &GoCalc{dataStack: NewStack(DATA_STACK_INIT_SIZE), program: make([]Instruction, 100), ip: 0, mu: &sync.Mutex{}}, nil
}

// GetIP returns with the actual value of the Instruction Pointer
func (c GoCalc) GetIP() int {
	return c.ip
}

// GetIP returns with the actual value of the Stack Pointer of the Data Stack
func (c GoCalc) GetDataStackPointer() int {
	return c.dataStack.GetPointer()
}

// Compile parses the `source` code, and generates the corresponding sequence of instructions into the `program` store
func (c GoCalc) Compile(source string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	// Make sure to have an end of an empty program
	c.program[0] = nil

	// Parse the source code
	parseResultAST, _ := Parse(source)

	// Encode the parsed results into a series of instructions
	fmt.Fprintf(debugWriter, "\n# source code: %s\n", source)
	fmt.Fprintf(debugWriter, "# program code:\n")
	nextInstruction := c.encode(parseResultAST, 0)
	c.program[nextInstruction] = nil
	fmt.Fprintf(debugWriter, "NIL\n\n")
}

// encode converts a node from the parsing results into an intruction, that also places ito the program.
func (c GoCalc) encode(node parc.Result, nextInstruction int) int {
	switch n := node.(type) {
	case Term:
		fmt.Fprintf(debugWriter, "# Term:\n")
		nextInstruction = c.encode(n.Operand_A, nextInstruction)
		nextInstruction = c.encode(n.Operand_B, nextInstruction)
		nextInstruction = c.encode(n.Operator, nextInstruction)

	case Expression:
		fmt.Fprintf(debugWriter, "# Expression:\n")
		nextInstruction = c.encode(n.Operand_A, nextInstruction)
		nextInstruction = c.encode(n.Operand_B, nextInstruction)
		nextInstruction = c.encode(n.Operator, nextInstruction)

	case Operator:
		switch n.Value {
		case "+":
			c.program[nextInstruction] = add
			fmt.Fprintf(debugWriter, "ADD\n")
		case "-":
			c.program[nextInstruction] = sub
			fmt.Fprintf(debugWriter, "SUB\n")
		case "/":
			c.program[nextInstruction] = div
			fmt.Fprintf(debugWriter, "DIV\n")
		case "*":
			c.program[nextInstruction] = mul
			fmt.Fprintf(debugWriter, "MUL\n")
		}
		nextInstruction++

	case Number:
		c.program[nextInstruction] = literal(&(c.dataStack), n.Value)
		fmt.Fprintf(debugWriter, "LIT %v\n", n.Value)
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
		//fmt.Printf("IP: %d, DSP: %d\n", c.GetIP(), c.dataStack.GetPointer())
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
