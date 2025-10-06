package gocalc

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/tombenke/go-12f-common/must"
	"github.com/tombenke/parc"
	"io"
	"os"
)

// PrintGraph generates a graphviz dot file from the AST tree, that is the result of the parser
func PrintDiagram(ast parc.Result, machineCode []string, title, fileName string) {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, `graph {
labelloc=t
label=<source code: %s <BR ALIGN="LEFT"/>>
fontname="Helvetica,Arial,sans-serif"
node [fontname="Helvetica,Arial,sans-serif" margin=0.1 fontcolor=blue fontsize=8 width=0.5 style=filled fillcolor=white color=blue]
edge [fontname="Helvetica,Arial,sans-serif"]
subgraph cluster_0 {
    label="syntax tree"

`, title)

	// Print AST generates a graphviz dot file from the AST tree, that is the result of the parser
	printASTNode(&buffer, ast, uuid.NewString())
	fmt.Fprintf(&buffer, `
}

subgraph cluster_1 {
    label="machine code"
`)

	// Print the generated machine code
	PrintMachineCode(&buffer, machineCode)

	fmt.Fprintf(&buffer, "\n}\n\n}")
	saveBytesToFile(buffer.Bytes(), fileName)
}

func PrintMachineCode(out io.Writer, machineCode []string) {
	fmt.Fprintf(out, `    machine_code [label=<`)
	for _, code := range machineCode {
		fmt.Fprintf(out, `%s<BR ALIGN="LEFT"/>`, code)
	}
	fmt.Fprintf(out, `> fontname="Courier:bold" shape=none fontcolor=black fontsize=14 nojustify=true];`)
}

// printAST generates the dot format graph representation of the selected AST node and its children
func printASTNode(out io.Writer, node parc.Result, nodeID string) {
	switch n := node.(type) {
	case Term:
		op_A_ID := uuid.NewString()
		operation_ID := uuid.NewString()
		op_B_ID := uuid.NewString()
		fmt.Fprintf(out, "    \"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, op_A_ID)
		printASTNode(out, n.Operand_A, op_A_ID)
		printASTNode(out, n.Operator, operation_ID)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, operation_ID)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, op_B_ID)
		printASTNode(out, n.Operand_B, op_B_ID)

	case Expression:
		op_A_ID := uuid.NewString()
		operation_ID := uuid.NewString()
		op_B_ID := uuid.NewString()
		fmt.Fprintf(out, "    \"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, op_A_ID)
		printASTNode(out, n.Operand_A, op_A_ID)
		printASTNode(out, n.Operator, operation_ID)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, operation_ID)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, op_B_ID)
		printASTNode(out, n.Operand_B, op_B_ID)

	case Operator:
		operator_ID := uuid.NewString()
		fmt.Fprintf(out, "    \"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "    \"%s\" [label=\"%s\" color=red fontcolor=red fontsize=12]\n", operator_ID, n.Value)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, operator_ID)

	case Number:
		operand_ID := uuid.NewString()
		fmt.Fprintf(out, "    \"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "    \"%s\" [label=\"%.2f\" color=green fontcolor=green fontsize=12]\n", operand_ID, n.Value)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, operand_ID)

	case Constant:
		operand_ID := uuid.NewString()
		fmt.Fprintf(out, "    \"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "    \"%s\" [label=\"%s\" color=green fontcolor=green fontsize=12]\n", operand_ID, n.Name)
		fmt.Fprintf(out, "    \"%s\" -- \"%s\"\n", nodeID, operand_ID)
	}
}

// saveBytesToFile saves the bytes data object into the file
func saveBytesToFile(data []byte, fileName string) {
	// Create the output file
	file := must.MustVal(os.Create(fileName))
	defer func() { must.Must(file.Close()) }()

	// Write data into the file
	must.MustVal(file.Write(data))
}
