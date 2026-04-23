package main

func sumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}
