package main

import "fmt"

func main() {

	value := []int{1, 2, 3}

	addWrong(value)
	fmt.Println(value)
	fmt.Println(len(value))

	value = addRight(value)
	fmt.Println(value)
}

func addWrong(nums []int) {
	nums = append(nums, 4)
}

func addRight(nums []int) []int {
	return append(nums, 4)
}
