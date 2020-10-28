package main

import "testing"

func TestAdd(t *testing.T) {

	var inventory Inventory

	testCases := []struct {
		product      Product
		mapLenth     int
		returnsError bool
	}{
		{
			Product{1, "product1"}, 1, false,
		},
		{
			Product{0, "product2"}, 1, true,
		},
		{
			Product{2, "product3"}, 2, false,
		},
		{
			Product{2, "product4"}, 2, true,
		},
	}

	for i, testCase := range testCases {

		err := inventory.Add(testCase.product)
		amountOfProducts := len(inventory.products)

		if err != nil && !testCase.returnsError {

			t.Errorf("TestAdd, case %d: shouldn't return an error\n", i)

		}

		if err == nil && testCase.returnsError {

			t.Errorf("TestAdd, case %d: should return an error\n", i)

		}

		if amountOfProducts != testCase.mapLenth {

			t.Errorf("TestAdd, case %d: the amount of products should be %d instead of %d\n", i, testCase.mapLenth, amountOfProducts)

		}

	}

}
