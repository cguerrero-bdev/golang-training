package different_asserts

import (
	"errors"
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
