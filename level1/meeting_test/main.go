package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (stack *Stack) Push(i int) {

	if stack == nil {
		stack.items = make([]int, 0)
	}

	stack.items = append(stack.items, i)

}

func (stack *Stack) Pop() (int, error) {

	if stack.items == nil || len(stack.items) == 0 {
		return 0, errors.New("The stack is empty")
	}

	result := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]
	return result, nil

}

func (stack *Stack) Peek() (int, error) {

	if stack.items == nil || len(stack.items) == 0 {
		return 0, errors.New("The stack is empty")
	}

	result := stack.items[len(stack.items)-1]
	return result, nil

}

func (stack *Stack) IsEmpty() bool {

	return stack.items == nil || len(stack.items) == 0

}

func main() {

	stack := Stack{}

	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())

	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack)

	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack)
	fmt.Println(stack.IsEmpty())

}
