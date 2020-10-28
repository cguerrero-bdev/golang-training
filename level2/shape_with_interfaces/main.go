package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length, width float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.length * rectangle.width
}

func (rectangle *Rectangle) Perimeter() float32 {
	return (rectangle.length * 2) + (rectangle.width * 2)
}

type Circle struct {
	radius float32
}

func (circle *Circle) Area() float32 {

	return math.Pi * float32(math.Sqrt(float64(circle.radius)))
}

func (circle *Circle) Perimeter() float32 {
	return 2 * math.Pi * circle.radius
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

func PrintShapeInf(shape Shape) {

	fmt.Printf("Area: %f\n", shape.Area())
	fmt.Printf("Perimeter: %f\n", shape.Area())
}

func main() {

	fmt.Println("* Rectangle Info *")
	PrintShapeInf(&Rectangle{10, 20})

	fmt.Println("* Circle Info *")
	PrintShapeInf(&Circle{20})
}
