package gocalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Create and startup a GoCalc component, then compile and run all the test cases with the Run() method
func TestGoCalcRun(t *testing.T) {
	calc, err := NewGoCalc()
	assert.Nil(t, err)
	for tcIdx, testCase := range testCases {
		calc.Compile(testCase.Formula)
		result := *calc.Run()
		assert.Equal(t, testCase.Result, float64(result))

		PrintDiagram(calc.GetAST(), calc.GetProgramDebug(), testCase.Formula, fmt.Sprintf("../docs/tc_%d.dot", tcIdx))
	}
}

// Create and startup a GoCalc component, then compile and run all the test cases with the RunAST() method
func TestGoCalcRunAST(t *testing.T) {
	calc, err := NewGoCalc()
	assert.Nil(t, err)
	for _, testCase := range testCases {
		calc.Compile(testCase.Formula)
		result := *calc.RunAST()
		assert.Equal(t, testCase.Result, float64(result))
	}
}

// Do benchmark on compile time
func BenchmarkGoCalcCompile(b *testing.B) {
	calc, _ := NewGoCalc()
	b.StartTimer()
	calc.Compile(benchmarkFormula)
	b.StopTimer()
	calc.Run()
}

// Do benchmark on runtime of compiled code
func BenchmarkGoCalcRun(b *testing.B) {
	calc, _ := NewGoCalc()
	calc.Compile(benchmarkFormula)
	b.ResetTimer()
	calc.Run()
}

// Do benchmark on runtime of compiled code
func BenchmarkGoCalcRunAST(b *testing.B) {
	calc, _ := NewGoCalc()
	calc.Compile(benchmarkFormula)
	b.ResetTimer()
	calc.RunAST()
}
