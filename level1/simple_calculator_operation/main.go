package main

import (
	"errors"
	"fmt"
)

type Calculator struct {
	result float32
}

func (c *Calculator) Add(number float32) float32 {

	c.result += number
	return c.result
}

func (c *Calculator) Subtract(number float32) float32 {

	c.result -= number
	return c.result
}

func (c *Calculator) Multiply(number float32) float32 {

	c.result *= number
	return c.result
}

func (c *Calculator) Divide(number float32) (float32, error) {

	if number == 0 {
		return 0, errors.New("the divisor cannot be 0")
	}

	c.result /= number
	return c.result, nil
}

func main() {

	calculator := Calculator{0}
	fmt.Println(calculator.Add(30))
	fmt.Println(calculator.Subtract(10))
	fmt.Println(calculator.Multiply(5))

	result, err := calculator.Divide(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = calculator.Divide(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
