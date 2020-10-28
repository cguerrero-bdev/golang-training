package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Printf("1")
			time.Sleep(2000)
		}
	}()

	for i := 0; i < 1000; i++ {
		fmt.Printf("0")
		time.Sleep(2000)
	}

}
