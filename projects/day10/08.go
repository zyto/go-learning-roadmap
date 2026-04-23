package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 20))
	fmt.Println(sumAll(1, 2, 3, 4))
	fmt.Println(sumAll())

	values := []int{5, 10, 15}
	fmt.Println(sumAll(values...))
}

func sumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}
