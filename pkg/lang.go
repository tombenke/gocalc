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

///type Operation struct {
///	Tag       string
///	Operation string
///	Operand_A parc.Result
///	Operand_B parc.Result
///}

type Atom struct {
	Tag   string
	Value parc.Result
}

type Term struct {
	Tag       string
	Operand_A parc.Result
	Operand_B parc.Result
	Operator  Operator
}

type Expression struct {
	Tag       string
	Operand_A parc.Result
	Operand_B parc.Result
	Operator  Operator
}

type Number struct {
	Tag   string
	Value StackData
}

type Operator struct {
	Tag   string
	Value string
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
	var atom parc.Parser
	var term parc.Parser
	var expression parc.Parser
	var number parc.Parser
	var realNumber parc.Parser
	var intNumber parc.Parser
	var mulOperator parc.Parser
	var addOperator parc.Parser

	number = *parc.Choice(&realNumber, &intNumber)

	intNumber = *parc.Map(parc.Integer, func(in parc.Result) parc.Result {
		operand := Number{
			Tag:   "NUMBER",
			Value: StackData(in.(int)),
		}
		return parc.Result(operand)
	})

	realNumber = *parc.Map(parc.RealNumber, func(in parc.Result) parc.Result {
		operand := Number{
			Tag:   "NUMBER",
			Value: StackData(in.(float64)),
		}
		return parc.Result(operand)
	})

	mulOperator = *parc.Map(parc.Choice(parc.Str("*"), parc.Str("/")), func(in parc.Result) parc.Result {
		return parc.Result(Operator{
			Tag:   "OPERATOR",
			Value: in.(string),
		})
	})

	addOperator = *parc.Map(parc.Choice(parc.Str("+"), parc.Str("-")), func(in parc.Result) parc.Result {
		return parc.Result(Operator{
			Tag:   "OPERATOR",
			Value: in.(string),
		})
	})

	atom = *parc.Map(parc.Choice(&number, parc.SequenceOf(
		parc.Str("("),
		parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
		&expression,
		parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
		parc.Str(")"),
	)), func(in parc.Result) parc.Result {
		var result parc.Result
		switch i := in.(type) {
		case Number:
			result = parc.Result(i)
		case []parc.Result:
			result = parc.Result(i[2])
		}
		return result
	})

	term = *parc.Map(parc.SequenceOf(
		&atom,
		parc.ZeroOrMore(
			parc.SequenceOf(
				parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
				&mulOperator,
				parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
				&atom,
			),
		),
	), func(in parc.Result) parc.Result {
		arr := in.([]parc.Result)
		arr1 := arr[1].([]parc.Result)
		if len(arr1) == 0 {
			return parc.Result(arr[0])
		}
		arr11 := arr1[0].([]parc.Result)
		return parc.Result(Term{Tag: "TERM", Operand_A: arr[0], Operator: arr11[1].(Operator), Operand_B: arr11[3]})
	})

	expression = *parc.Map(parc.SequenceOf(
		&term,
		parc.ZeroOrMore(
			parc.SequenceOf(
				parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
				&addOperator,
				parc.ZeroOrMore(parc.Cond(parc.IsWhitespace)),
				&term,
			),
		),
	), func(in parc.Result) parc.Result {
		arr := in.([]parc.Result)
		arr1 := arr[1].([]parc.Result)
		if len(arr1) == 0 {
			return parc.Result(arr[0])
		}
		arr11 := arr1[0].([]parc.Result)
		return parc.Result(Expression{Tag: "EXPRESSION", Operand_A: arr[0], Operator: arr11[1].(Operator), Operand_B: arr11[3]})
	})

	return expression
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
	case Term:
		op_A_ID := uuid.NewString()
		operation_ID := uuid.NewString()
		op_B_ID := uuid.NewString()
		fmt.Fprintf(out, "\"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_A_ID)
		printASTNode(out, n.Operand_A, op_A_ID)
		printASTNode(out, n.Operator, operation_ID)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, operation_ID)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_B_ID)
		printASTNode(out, n.Operand_B, op_B_ID)

	case Expression:
		op_A_ID := uuid.NewString()
		operation_ID := uuid.NewString()
		op_B_ID := uuid.NewString()
		fmt.Fprintf(out, "\"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_A_ID)
		printASTNode(out, n.Operand_A, op_A_ID)
		printASTNode(out, n.Operator, operation_ID)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, operation_ID)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, op_B_ID)
		printASTNode(out, n.Operand_B, op_B_ID)

	case Operator:
		operator_ID := uuid.NewString()
		fmt.Fprintf(out, "\"%s\" [label=\" %s \" shape=box]\n", nodeID, n.Tag)
		fmt.Fprintf(out, "\"%s\" [label=\"%s\" color=red fontcolor=red fontsize=12]\n", operator_ID, n.Value)
		fmt.Fprintf(out, "\"%s\" -- \"%s\"\n", nodeID, operator_ID)

	case Number:
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
