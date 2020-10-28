package main

import "testing"

func TestEnqueue(t *testing.T) {

	var testCases = []struct {
		item         int
		queueContent []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
	}

	var queue Queue

	for i, testCase := range testCases {

		Enqueue(&queue, testCase.item)

		if !areEqualSlides(testCase.queueContent, queue.items) {
			t.Errorf("Test Case %d error ", i)
		}
	}

}

func TestDequeue(t *testing.T) {

	var testCases = []struct {
		want         int
		queueContent []int
		err          bool
	}{
		{1, []int{2, 3}, false},
		{2, []int{3}, false},
		{3, []int{}, false},
		{0, []int{}, true},
	}

	queue := Queue{[]int{1, 2, 3}}

	for i, testCase := range testCases {

		item, err := Dequeue(&queue)

		isError := item != testCase.want
		isError = isError || !areEqualSlides(testCase.queueContent, queue.items)
		isError = isError || (err != nil) != testCase.err

		if isError {
			t.Errorf("Test Case %d error ", i)
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
