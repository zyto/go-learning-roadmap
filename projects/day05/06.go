package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	slice1 := nums[1:4]
	slice2 := nums[2:5]
	slice1[1] = 99

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(nums)
}
