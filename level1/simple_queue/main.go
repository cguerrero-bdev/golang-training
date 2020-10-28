package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

func Enqueue(queue *Queue, item int) {

	if queue.items == nil {
		queue.items = make([]int, 0)
	}

	queue.items = append(queue.items, item)
}

func Dequeue(queue *Queue) (int, error) {

	if queue.items == nil || len(queue.items) == 0 {
		return 0, errors.New("Empty queue")
	}
	result := queue.items[0]
	queue.items = queue.items[1:]
	return result, nil
}

func main() {

	queue := &Queue{}

	fmt.Println(Dequeue(queue))

	Enqueue(queue, 1)
	Enqueue(queue, 2)
	Enqueue(queue, 3)

	fmt.Println(queue)
	fmt.Println(Dequeue(queue))
	fmt.Println(queue)

	fmt.Println(Dequeue(queue))
	fmt.Println(queue)

	fmt.Println(Dequeue(queue))
	fmt.Println(queue)

	fmt.Println(Dequeue(queue))
	fmt.Println(queue)
}
