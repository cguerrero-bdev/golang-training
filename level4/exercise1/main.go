package main

import (
	"fmt"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("Hello World\n")
		}
	}()

	for i := 0; i < 100; i++ {
		fmt.Printf("Main function\n")
	}

}
