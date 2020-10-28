package main

import "testing"

func TestPrintInterfaceInformation(t *testing.T) {

	testCases := []struct {
		customInterface customInterface
		want            string
	}{
		{&interfaceImplementor{}, "Implemented by interfaceImplementor"},
		{&interfaceImplementor2{}, "Implemented by interfaceImplementor2"},
		{&interfaceImplementor3{}, "Implemented by interfaceImplementor3"},
	}

	for i, testCase := range testCases {
		if printInterfaceInformation(testCase.customInterface) != testCase.want {
			t.Errorf("Test case %d: incorrect message", i+1)
		}
	}
}
