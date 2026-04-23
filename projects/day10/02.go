package main

import "fmt"

func main() {
	fmt.Println(sumSlice([]int{4, 7, 1, 8}))
}

func sumSlice(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}
