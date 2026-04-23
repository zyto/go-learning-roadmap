package main

import (
	"fmt"
)

func main() {
	var s1 []int
	printResult(s1)

	s2 := []int{99, 0, -5}
	printResult(s2)
}

func minMax(nums []int) (int, int, bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	min, max := nums[0], nums[0]

	for _, n := range nums[1:] {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max, true
}

func printResult(s []int) {
	min, max, ok := minMax(s)

	if !ok {
		fmt.Println("Пустой срез!")
	} else {
		fmt.Printf("Минимум: %d, Максимум: %d\n", min, max)
	}
}
