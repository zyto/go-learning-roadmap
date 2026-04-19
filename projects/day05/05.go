package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 3, 5)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(1000)
	}

	fmt.Println(nums)

	nums = append(nums, 77)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}
