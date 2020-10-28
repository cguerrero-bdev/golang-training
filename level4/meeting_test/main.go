package main

import "fmt"

func findSubstring(sList []string, s2 string, channel chan string) {

	for i := 0; i < len(sList); i++ {

		if sList[i] == s2 {
			channel <- sList[i]
		}
	}

}

func main() {

	stringList := []string{"test", "exam", "test", "exam"}

	channel := make(chan string, len(stringList))

	go findSubstring(stringList[0:1], "test", channel)

	//go findSubstring(stringList[2:3], "test", channel)

	for range channel {
		result, ok := <-channel
		if ok {
			fmt.Println(result)
		}
	}

}
