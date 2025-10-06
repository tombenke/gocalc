package gocalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangParser(t *testing.T) {
	parser := buildParser()

	for _, testCase := range testCases {
		input := testCase.Formula
		parseResults := parser.Parse(&input)
		assert.NotNil(t, parseResults.Results)
		assert.Equal(t, testCase.Error, parseResults.Err)
	}
}
