package main

import "fmt"

func main() {
	nums := make([]int, 2, 5)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	nums[0] = 10
	nums[1] = 20

	nums = append(nums, 30, 40)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}
