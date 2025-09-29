package gocalc

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangParse(t *testing.T) {
	input := "((10 * 2) + ((50 / 3) - 2))"
	parser := buildParser()
	parseResults := parser.Parse(&input)
	assert.NotNil(t, parseResults)
	//fmt.Printf("%s => interpreter => %+v\n", input, parseResults)
}
