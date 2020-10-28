package main

import "fmt"

func valueProducer(channel chan int) {
	channel  <- 888
}

func main() {

	channel := make(chan int)

	go valueProducer(channel)

	v, ok := <-channel
	fmt.Printf("Value from the channel: %d", <-channel)
}
