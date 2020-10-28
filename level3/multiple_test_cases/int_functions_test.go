package int_functions

import "testing"

var table = []struct {
	n1, n2, want int
}{
	{1, 5, 1},
	{6, 5, 5},
	{25, 5, 5},
}

func TestTableIntMin1(t *testing.T) {

	for _, test := range table {
		if result := IntMin(test.n1, test.n2); result != test.want {

			t.Errorf("IntMin(%d, %d) = %d\nExpected result: %d", test.n1, test.n2, test.want, result)
		}
	}
}
