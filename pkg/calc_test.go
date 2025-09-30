package gocalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCases = []TestCase{
		TestCase{Formula: "42", Result: 42., Error: nil},
		//TestCase{Formula: "((10.*2) + ((60.0 / 3) - 2.))", Result: 38., Error: nil},
		TestCase{Formula: "((10.* \t \n\n 2) +((60.0/3)-2.))", Result: 38., Error: nil},
	}
	benchmarkFormula = "((10. * 2) + ((60.0 / 3) - 2.))"
)

type TestCase struct {
	Formula string
	Result  float64
	Error   error
}

// Create and startup a GoCalc component, then compile and run all the test cases
func TestGoCalc(t *testing.T) {
	calc, err := NewGoCalc()
	assert.Nil(t, err)
	for _, testCase := range testCases {
		calc.Compile(testCase.Formula)
		//fmt.Printf("calc: %+v\nIP: %d, DSP: %d\n\n", calc, calc.GetIP(), calc.GetDataStackPointer())
		result := *calc.Run()
		assert.Equal(t, testCase.Result, float64(result))
		fmt.Printf("IP: %d, DSP: %d\nresult: %.2f\n\n", calc.GetIP(), calc.GetDataStackPointer(), result)
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
