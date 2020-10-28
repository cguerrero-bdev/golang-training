package main

import "testing"

func TestArea(t *testing.T) {

	testCases := []struct {
		shape Shape
		area  float32
	}{
		{&Rectangle{10.0, 2}, 20.0},
		{&Circle{20.0}, 14.049630},
		{&Rectangle{15.0, 3}, 45.0},
		{&Circle{2.0}, 4.442883},
	}

	for i, testCase := range testCases {
		if a := testCase.shape.Area(); a != testCase.area {
			t.Errorf("Test case %d. expected area: %f, returned area: %f", i, testCase.area, a)
		}
	}

}

func TestPerimeter(t *testing.T) {

	testCases := []struct {
		shape     Shape
		perimeter float32
	}{
		{&Rectangle{10.0, 2}, 24.0},
		{&Circle{20.0}, 125.663712},
		{&Rectangle{15.0, 3}, 36},
		{&Circle{2.0}, 12.566371},
	}

	for i, testCase := range testCases {
		if p := testCase.shape.Perimeter(); p != testCase.perimeter {
			t.Errorf("Test case %d. expected area: %f, returned area: %f", i+1, testCase.perimeter, p)
		}
	}

}
