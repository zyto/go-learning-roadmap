package main

import "fmt"

func main() {
	nums := []int{3, -1, 5, -7}
	extra := 10

	replaceNegativeWithZero(nums)

	fmt.Println(nums)

	nums = appendIfPositive(nums, extra)
	fmt.Println(nums)

	fmt.Println(resetIfNil(&nums[0]))
}

func replaceNegativeWithZero(nums []int) {
	for i, v := range nums {
		if v < 0 {
			nums[i] = 0
		}
	}
}

func appendIfPositive(nums []int, value int) []int {
	if value > 0 {
		nums = append(nums, value)
	}

	return nums
}

func resetIfNil(p *int) int {
	if p == nil {
		return 0
	}

	return *p
}
