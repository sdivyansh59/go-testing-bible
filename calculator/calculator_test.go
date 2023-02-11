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
	testCase := TestCase {
		value: 371,
		expected: true,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}


func TestNegativeCalculeIsArmstrong(t *testing.T) {
	testCase := TestCase {
		value: 350,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}