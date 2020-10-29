package main

import "testing"

func TestApplyDiscount(t *testing.T) {

	testCases := []struct {
		product      Product
		productPrice float32
	}{
		{&Book{"Any Book", 20}, 18},
		{&Game{"Any Game", 20}, 16},
		{&Book{"Go Manual", 40}, 36},
		{&Game{"Monopoly", 40}, 32},
	}

	for i, testCase := range testCases {

		applyDiscount(testCase.product)
		if testCase.product.GetPrice() != testCase.productPrice {
			t.Errorf("Test case %d. expected price %f, returned price %f", i+1, testCase.product.GetPrice(), testCase.productPrice)
		}

	}

}
