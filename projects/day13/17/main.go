package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	p := &nums[1]

	*p = 99

	fmt.Println(nums)
}
