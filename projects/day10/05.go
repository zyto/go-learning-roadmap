package main

import "fmt"

func main() {
	nums := []int{3, 10, 7, 15, 2}
	limit := 7

	fmt.Println(filterGreaterThan(nums, limit))
}

func filterGreaterThan(nums []int, limit int) []int {
	s := make([]int, 0, len(nums))

	for _, v := range nums {
		if v > limit {
			s = append(s, v)
		}
	}

	return s
}
