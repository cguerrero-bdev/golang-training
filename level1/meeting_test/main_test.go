package main

import "testing"

func TestPush(t *testing.T) {

	testCases := []struct {
		newItem        int
		resultingItems []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
	}

	var stack Stack

	for i, testCase := range testCases {

		stack.Push(testCase.newItem)

		if !areEqualSlides(stack.items, testCase.resultingItems) {
			t.Errorf("Test case %d. Values in the stack %v, expected values %v ", i, testCase.resultingItems, stack.items)
		}
	}
}

func TestPop(t *testing.T) {

	testCases := []struct {
		want           int
		resultingItems []int
		err            bool
	}{

		{3, []int{1, 2}, false},
		{2, []int{1}, false},
		{1, []int{}, false},
		{0, []int{}, true},
	}

	var stack = Stack{[]int{1, 2, 3}}

	for i, testCase := range testCases {

		result, err := stack.Pop()

		if testCase.err && err == nil {
			t.Errorf("Test case %d. error expected ", i)
		}

		if !testCase.err && err != nil {
			t.Errorf("Test case %d. error not expected ", i)
		}

		if result != testCase.want {
			t.Errorf("Test case %d. Wrong return value %d, expected values %d ", i, result, testCase.want)
		}
		if !areEqualSlides(stack.items, testCase.resultingItems) {
			t.Errorf("Test case %d. Values in the stack %v, expected values %v ", i, testCase.resultingItems, stack.items)
		}
	}
}

func areEqualSlides(slide1 []int, slide2 []int) bool {
	if len(slide1) != len(slide2) {
		return false
	}

	for i, item := range slide1 {
		if item != slide2[i] {
			return false
		}
	}

	return true
}
