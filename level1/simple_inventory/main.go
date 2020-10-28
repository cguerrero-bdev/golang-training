package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID   int
	Name string
}

type Inventory struct {
	products   map[string]*Product
	productsId map[int]*Product
}

func (inventory *Inventory) Add(product Product) error {

	if product.ID < 1 {
		return errors.New("ID can not be empty")
	}

	if inventory.productsId != nil && inventory.productsId[product.ID] != nil {
		return errors.New("ID can not be duplicated")
	}

	if inventory.products == nil {
		inventory.products = make(map[string]*Product)
		inventory.productsId = make(map[int]*Product)
	}

	inventory.products[product.Name] = &product
	inventory.productsId[product.ID] = &product

	return nil
}

func main() {

	inventory := Inventory{}

	err := inventory.Add(Product{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inventory.products)
	}

	err = inventory.Add(Product{1, "Product1"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inventory.products)
	}

	err = inventory.Add(Product{1, "Product1"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inventory.products)
	}

	err = inventory.Add(Product{2, "Product2"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inventory.products)
	}

	err = inventory.Add(Product{3, "Product3"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inventory.products)
	}

}
