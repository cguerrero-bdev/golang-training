package main

import "fmt"

type Product interface {
	PrintInformation()
	Discount(radio float32)
	GetPrice() float32
}

type Game struct {
	name  string
	price float32
}

func (game *Game) PrintInformation() {
	fmt.Printf("Name: %s, Price %f\n", game.name, game.price)
}

func (game *Game) Discount(radio float32) {
	game.price -= (game.price * radio / 100)
}

func (game *Game) GetPrice() float32 {
	return game.price
}

type Book struct {
	name  string
	price float32
}

func (book *Book) PrintInformation() {
	fmt.Printf("Name: %s, Price %f\n", book.name, book.price)
}

func (game *Book) GetPrice() float32 {
	return game.price
}

func (book *Book) Discount(radio float32) {
	book.price -= (book.price * radio / 100)
}

func applyDiscount(product Product) {

	switch product.(type) {

	case *Book:
		product.Discount(10)
	case *Game:
		product.Discount(20)

	}

	product.PrintInformation()
}

func main() {
	applyDiscount(&Book{"Any Book", 20})
	applyDiscount(&Game{"Any Game", 20})
}
