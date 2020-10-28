package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var testCases = []struct {
	want  []int
	param []int
	age1  int
	age2  int
}{
	{[]int{40, 45}, []int{40, 45}, 20, 50},
	{[]int{30, 25, 40}, []int{30, 17, 25, 40}, 20, 50},
	{[]int{30, 40, 50}, []int{17, 30, 40, 50, 60}, 18, 55},
}

func TestTableFilter(t *testing.T) {

	for i := 0; i < len(testCases); i++ {

		if !Equal(testCases[i].want, filter(testCases[i].age1, testCases[i].age2, testCases[i].param)) {
			t.Errorf("It's not the righ filter in case %d", i+1)
		}

	}
}
