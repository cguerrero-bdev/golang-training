package int_functions

import "testing"

func TestIntMin1(t *testing.T) {

	result := IntMin(350, 50)

	if result != 50 {

		t.Errorf("IntMin(350, 50) = %d\nExpected result: 50", result)
	}
}
