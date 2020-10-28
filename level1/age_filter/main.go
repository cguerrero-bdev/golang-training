package main

import "fmt"

func filter(age1 int, age2 int, ages []int) []int {

	result := make([]int, 0)

	for i := 0; i < len(ages); i++ {

		if ages[i] > age1 && ages[i] < age2 {
			result = append(result, ages[i])
		}

	}

	return result
}

func main() {

	ages := []int{30, 17, 25, 40, 70, 45}
	fmt.Println(ages)
	fmt.Println(filter(30, 46, ages))

}
