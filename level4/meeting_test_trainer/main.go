package main

import (
	"fmt"
	"strings"
	"sync"
)

func findSubstr(list []string, match string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, s := range list {
		if strings.Contains(s, match) {
			ch <- s
		}
	}
}

func findSubstrInParallel(sList []string, match string) []string {

	result := make([]string, 0)
	outC := make(chan string)

	// consumer
	go func() {
		for s := range outC {
			result = append(result, s)
			fmt.Println(s)
		}
	}()
	// producers
	wg := sync.WaitGroup{}
	wg.Add(2)
	go findSubstr(sList[:(len(sList)/2)], match, outC, &wg)
	go findSubstr(sList[(len(sList)/2):], match, outC, &wg)
	wg.Wait()
	close(outC)

	return result
}

func main() {
	sList := []string{"abc", "abcd", "abcde", "abtest", "lallabcla", "abctest", "test", "atest"}
	outC := make(chan string)
	// consumer
	go func() {
		for s := range outC {
			fmt.Println(s)
		}
	}()
	// producers
	wg := sync.WaitGroup{}
	wg.Add(2)
	go findSubstr(sList[:(len(sList)/2)], "abc", outC, &wg)
	go findSubstr(sList[(len(sList)/2):], "abc", outC, &wg)
	wg.Wait()
	close(outC)
	// this will panic
	// outC <- "lala"
}
