package main

import "fmt"

func main() {
	nums := []int{3, 10, 2, 8, 5}
	limit := 5

	dst := filterByLimit(nums, &limit)
	fmt.Println(dst)
}

func filterByLimit(nums []int, limit *int) []int {
	if limit == nil {
		return nums
	}

	dst := make([]int, 0, len(nums))
	for _, n := range nums {
		if n <= *limit {
			dst = append(dst, n)
		}
	}

	return dst
}
