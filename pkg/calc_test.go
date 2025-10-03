package gocalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Create and startup a GoCalc component, then compile and run all the test cases
func TestGoCalc(t *testing.T) {
	calc, err := NewGoCalc()
	assert.Nil(t, err)
	for tcIdx, testCase := range testCases {
		calc.Compile(testCase.Formula)
		//fmt.Printf("calc: %+v\nIP: %d, DSP: %d\n\n", calc, calc.GetIP(), calc.GetDataStackPointer())
		result := *calc.Run()
		assert.Equal(t, testCase.Result, float64(result))
		fmt.Printf("IP: %d, DSP: %d\nresult: %.2f\n\n", calc.GetIP(), calc.GetDataStackPointer(), result)

		PrintDiagram(calc.GetAST(), calc.GetProgramDebug(), testCase.Formula, fmt.Sprintf("../docs/tc_%d.dot", tcIdx))
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
