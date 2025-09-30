package gocalc

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/tombenke/go-12f-common/must"
	"github.com/tombenke/parc"
	"io"
	"log"
	"os"
)

type Operation struct {
	Tag       string
	Operation string
	Operand_A parc.Result
	Operand_B parc.Result
}

type Operand struct {
	Tag   string
	Value StackData
}

// The parser instance
var parser parc.Parser

func init() {
	parser = buildParser()
}

// Parse uses the internal parser object that parses the input string and returns with the results in the form of AST tree.
func Parse(source string) (parc.Result, error) {

	parseResults := parser.Parse(&source)
	if parseResults.IsError {
		return nil, parseResults.Err
	}

	return parseResults.Results, nil
}

// buildParser creates a parser of the language. It is called by the init function.
func buildParser() parc.Parser {
	var expr parc.Parser
	var literal parc.Parser
	var realNumber parc.Parser
	var intNumber parc.Parser
	var operation parc.Parser

	operator := parc.Choice(parc.Str("+"), parc.Str("-"), parc.Str("*"), parc.Str("/"))

	literal = *parc.Choice(&realNumber, &intNumber)

	intNumber = *parc.Map(parc.Integer, func(in parc.Result) parc.Result {
		operand := Operand{
			Tag:   "LITERAL",
			Value: StackData(in.(int)),
		}
		return parc.Result(operand)
	})

	realNumber = *parc.Map(parc.RealNumber, func(in parc.Result) parc.Result {
		operand := Operand{
			Tag:   "LITERAL",
			Value: StackData(in.(float64)),
		}
		return parc.Result(operand)
	})

	expr = *parc.Choice(&literal, &operation)

	operation = *parc.Map(parc.SequenceOf(
		parc.Str("("),
		&expr,
		parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
		operator,
		parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
		&expr,
		parc.Str(")"),
	), func(in parc.Result) parc.Result {
		arr := in.([]parc.Result)
		op := Operation{
			Tag:       "OPERATION",
			Operation: arr[3].(string),
			Operand_A: arr[1],
			Operand_B: arr[5],
		}
		return parc.Result(op)
	})

	return expr
}

// PrintAST generates a graphviz dot file from the AST tree, that is the result of the parser
func PrintAST(ast parc.Result, title, fileName string) {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, `graph {
	label="%s"
    fontname="Helvetica,Arial,sans-serif"
    node [fontname="Helvetica,Arial,sans-serif" margin=0.1 fontcolor=blue fontsize=8 width=0.5 style=filled fillcolor=white color=blue]
    edge [fontname="Helvetica,Arial,sans-serif"]
	`, title)

	printASTNode(&buffer, ast, uuid.NewString())
	fmt.Fprintf(&buffer, "\n}")

	saveBytesToFile(buffer.Bytes(), fileName)
}

// printAST generates the dot format graph representation of the selected AST node and its children
func printASTNode(out io.Writer, node parc.Result, nodeID string) {
	switch n := node.(type) {
	case Operation:

		op_A_ID := uuid.NewString()
		operation_ID := uuid.NewString()
		op_B_ID := uuid.NewString()
		fmt.Fprintf(out, "\"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_A_ID)
		printASTNode(out, n.Operand_A, op_A_ID)
		fmt.Fprintf(out, "\"%s\" [label=\"%s\" color=red fontcolor=red fontsize=12]\n", operation_ID, n.Operation)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, operation_ID)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_B_ID)
		printASTNode(out, n.Operand_B, op_B_ID)

	case Operand:
		operand_ID := uuid.NewString()
		fmt.Fprintf(out, "\"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "\"%s\" [label=\"%.2f\" color=green fontcolor=green fontsize=12]\n", operand_ID, n.Value)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, operand_ID)
	}
}

// saveBytesToFile saves the bytes data object into the file
func saveBytesToFile(data []byte, fileName string) {
	// Create the output file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { must.Must(file.Close()) }()

	// Write data into the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error in write:", err)
		return
	}
}
