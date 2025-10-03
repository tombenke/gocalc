package gocalc

var (
	testCases = []TestCase{
		TestCase{Formula: "42", Result: 42., Error: nil},
		TestCase{Formula: "42+24", Result: 66., Error: nil},
		TestCase{Formula: "42+24*33", Result: 834., Error: nil},
		TestCase{Formula: "((10. * 2) + ((50.0 / 3.1415) - 2.))", Result: 33.915963711602735, Error: nil},
		TestCase{Formula: "((10.*2) + ((60.0 / 3) - 2.))", Result: 38., Error: nil},
		TestCase{Formula: "((10.* \t \n\n 2) +((60.0/3)-2.))", Result: 38., Error: nil},
		TestCase{Formula: "1 + 2 * 3", Result: 7., Error: nil},
		TestCase{Formula: "(1 + 2) * 3", Result: 9., Error: nil},
	}
	benchmarkFormula = testCases[4].Formula
)

type TestCase struct {
	Formula string
	Result  float64
	Error   error
}
