package main

import "fmt"

type customInterface interface {
	GetMessage() string
	GetValue() int
}

type interfaceImplementor struct{}

func (interfaceImplementor *interfaceImplementor) GetMessage() string {
	return "Implementor 1"
}

func (interfaceImplementor *interfaceImplementor) GetValue() int {
	return 1
}

type interfaceImplementor2 struct{}

func (interfaceImplementor *interfaceImplementor2) GetMessage() string {
	return "Implementor 2"
}

func (interfaceImplementor *interfaceImplementor2) GetValue() int {
	return 2
}

type interfaceImplementor3 struct{}

func (interfaceImplementor *interfaceImplementor3) GetMessage() string {
	return "Implementor 3"
}

func (interfaceImplementor *interfaceImplementor3) GetValue() int {
	return 3
}

func printInterfaceInformation(ci customInterface) {

	fmt.Printf("\nInterface message: %s\n", ci.GetMessage())
	fmt.Printf("Interface value: %d\n", ci.GetValue())

	switch ci.(type) {

	case *interfaceImplementor:
		fmt.Println("Implemented by interfaceImplementor")
	case *interfaceImplementor2:
		fmt.Println("Implemented by interfaceImplementor2")
	case *interfaceImplementor3:
		fmt.Println("Implemented by interfaceImplementor3")
	}
}

func main() {

	printInterfaceInformation(&interfaceImplementor{})

	printInterfaceInformation(&interfaceImplementor2{})

	printInterfaceInformation(&interfaceImplementor3{})
}
