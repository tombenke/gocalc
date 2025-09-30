package gocalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangParseEmpty(t *testing.T) {
	input := "42"
	parser := buildParser()
	parseResults := parser.Parse(&input)
	assert.NotNil(t, parseResults.Results)
	assert.Nil(t, parseResults.Err)
	fmt.Printf("%s => interpreter => %+v\n", input, parseResults)
}

func TestLangParse(t *testing.T) {
	input := "((10. * 2) + ((50.0 / 3.1415) - 2.))"
	parseResults, err := Parse(input)

	assert.NotNil(t, parseResults)
	assert.Nil(t, err)
	fmt.Printf("%s => interpreter => %+v\n", input, parseResults)
}

func TestPrintAST(t *testing.T) {
	input := "((10. * 2) + ((50.0 / 3.1415) - 2.))"
	parseResults, _ := Parse(input)
	PrintAST(parseResults, input, "../docs/ast.dot")
}
