package main

import "testing"

func TestAdd(t *testing.T) {

	var testCases = []struct {
		number float32
		want   float32
	}{
		{10.0, 10.0},
		{5.0, 15.0},
		{2.2, 17.2},
		{2.8, 20.0},
	}

	var calculator Calculator

	for _, testCase := range testCases {

		if calculator.Add(testCase.number) != testCase.want {
			t.Errorf("TestAdd: %f + %f should be: %f", calculator.result, testCase.number, testCase.want)
		}
	}
}

func TestDivide(t *testing.T) {

	testCases := []struct {
		number       float32
		want         float32
		returnsError bool
	}{
		{5.0, 20.0, false},
		{4.0, 5.0, false},
		{5.0, 1.0, false},
		{0.0, 0.0, true},
	}

	calculator := Calculator{100}

	for _, testCase := range testCases {

		result, err := calculator.Divide(testCase.number)
		if result != testCase.want && err != nil == testCase.returnsError {
			t.Errorf("TestAdd: %f divide by %f should be: %f", calculator.result, testCase.number, testCase.want)
		}
	}

}
