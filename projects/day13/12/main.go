package main

import "fmt"

func main() {
	value := []int{1, 2, 3}
	changeFirst(value)
	fmt.Println(value)
}

func changeFirst(nums []int) {
	if len(nums) == 0 {
		return
	}

	nums[0] = 100
}
