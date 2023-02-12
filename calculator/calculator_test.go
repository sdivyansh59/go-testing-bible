package calculator_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sdivyansh59/go-testing-bible/calculator"
) 

type TestCase struct {
	name string
	value int
	expected bool
	actual bool
}

func TestMain(m *testing.M) {
	fmt.Println("Hello World")
	ret := m.Run()
	fmt.Println("Test have executed")
	os.Exit(ret)
}

func TestCalculeIsArmstrong(t *testing.T) {
	t.Run("test for all 3 digit armstrong number", func(t *testing.T) {
		tests := []TestCase{
			TestCase{name: "Testing value for: 153",value: 153, expected: true},
			TestCase{name: "Testing value for: 370",value: 370, expected: true},
			TestCase{name: "Testing value for: 371",value: 371, expected: true},
			TestCase{name: "Testing value for: 407",value: 407, expected: true},	
		}

		for _,test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})	
	
}


func TestNegativeCalculeIsArmstrong(t *testing.T) {

	t.Run("should fail for caase 350", func(t *testing.T) {
		testCase := TestCase {
			value: 350,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should fail for caase 300", func(t *testing.T) {
		testCase := TestCase {
			value: 300,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
	
}



// Benchmark
//run: go test ./... -run=Benchmark -bench=.
func BenchmarkCalculateIsArmstrong370(b *testing.B){
	benchmarkCalculateIsArmstrong(370,b)
}

func BenchmarkCalculateIsArmstrong371(b *testing.B){
	benchmarkCalculateIsArmstrong(371,b)
}

func benchmarkCalculateIsArmstrong(input int , b * testing.B) {
	for i :=0 ; i< b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}