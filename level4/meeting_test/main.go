package main

import (
	"fmt"
)

func findString(sList []string, s2 string, i1 int, i2 int, channel chan int) {

	for i, iCounter := i1, 0; i <= i2; i++ {

		if sList[i] == s2 {
			channel <- i1 + iCounter
		}

		iCounter++
	}

	close(channel)

}

func main() {

	sList := []string{"one", "two", "three", "four", "one", "two", "three", "four"}
	stringToFind := "two"

	result := make([]int, 0)

	c := make(chan int, len(sList))
	findString(sList, stringToFind, 0, 3, c)

	c2 := make(chan int, len(sList))
	findString(sList, stringToFind, 4, 7, c2)

	for r := range c {
		result = append(result, r)
	}

	for r := range c2 {
		result = append(result, r)
	}

	fmt.Println(result)
}
