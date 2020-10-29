package main

import "testing"

func TestFindSubstr(t *testing.T) {

	testCases := []struct {
		stringList []string
		match      string
		result     []string
	}{
		{
			[]string{"abc", "abcd", "abcde", "abtest", "lallabcla", "abctest", "test", "atest"},
			"abc",
			[]string{"lallabcla", "abc", "abctest", "abcd", "abcde"},
		},
		{
			[]string{"one", "two", "three", "four", "one", "two", "three", "four"},
			"o",
			[]string{"one", "two", "one", "two", "four"},
		},
	}

	for i, testCase := range testCases {

		result := findSubstrInParallel(testCase.stringList, testCase.match)

		if !areEqualSlides(result, testCase.result) {

			t.Errorf("Test case %d. Expected value %v, returned value %v", i+1, testCase.result, result)
		}

	}
}

func areEqualSlides(slide1 []string, slide2 []string) bool {
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
