package gocalc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	formula = "((10 * 2) + ((60 / 3) - 2))"
)

// Create and startup a GoCalc component
func TestGoCalc(t *testing.T) {
	calc, err := NewGoCalc()
	assert.Nil(t, err)
	calc.Compile(formula)
	//fmt.Printf("calc: %+v\nIP: %d, DSP: %d\n\n", calc, calc.GetIP(), calc.GetDataStackPointer())
	result := *calc.Run()
	assert.Equal(t, 38., float64(result))
	fmt.Printf("IP: %d, DSP: %d\nresult: %.2f\n\n", calc.GetIP(), calc.GetDataStackPointer(), result)
}

func BenchmarkGoCalcCompile(b *testing.B) {
	calc, _ := NewGoCalc()
	b.StartTimer()
	calc.Compile(formula)
	b.StopTimer()
	calc.Run()
}

func BenchmarkGoCalcRun(b *testing.B) {
	calc, _ := NewGoCalc()
	calc.Compile(formula)
	b.ResetTimer()
	calc.Run()
}
