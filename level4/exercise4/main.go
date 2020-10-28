package main

import (
	"fmt"
	"sync"
	"time"
)

func increase(routineName string, mutex *sync.Mutex) {

	//	mutex.Lock()

	for i := 0; i < 10; i++ {
		fmt.Printf("%s: new value %d\n", routineName, i)

	}

	//	defer mutex.Unlock()

}

func main() {

	var mutex sync.Mutex
	go increase("Routine 2", &mutex)
	increase("Routine 1", &mutex)

	time.Sleep(time.Second * 5)
}
