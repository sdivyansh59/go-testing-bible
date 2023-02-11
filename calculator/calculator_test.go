package calculator_test

import (
	"testing"

	"github.com/sdivyansh59/go-testing-bible/calculator"
) 

type TestCase struct {
	value int
	expected bool
	actual bool
}

func TestCalculeIsArmstrong(t *testing.T) {

	t.Run("should return true for 371", func(t *testing.T) {
		testCase := TestCase {
			value: 371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should return true for 370", func(t *testing.T) {
		testCase := TestCase {
			value: 370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
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