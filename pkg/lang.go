package gocalc

import (
	"github.com/tombenke/parc"
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

func buildParser() parc.Parser {
	var expr parc.Parser
	var operation parc.Parser

	operator := parc.Choice(parc.Str("+"), parc.Str("-"), parc.Str("*"), parc.Str("/"))

	expr = *parc.Choice(
		parc.Map(parc.Integer, func(in parc.Result) parc.Result {
			operand := Operand{
				Tag:   "LITERAL",
				Value: StackData(in.(int)),
			}
			return parc.Result(operand)
		}),
		&operation,
	)

	operation = *parc.Map(parc.SequenceOf(
		parc.Str("("),
		&expr,
		parc.Str(" "),
		operator,
		parc.Str(" "),
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
