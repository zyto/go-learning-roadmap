package main

import "fmt"

func main() {
	nums := []int{12, 5, 18, 3, 25, 8, 30}

	fmt.Printf("Сумма всех чисел: %d\n", sumSlice(nums))
	slice10 := filterGreaterThan(nums, 10)
	fmt.Printf("Все числа больше 10: %v\n", slice10)

	greater20, ok := findFirstGreaterThan(nums, 20)

	if ok {
		fmt.Printf("Первое число больше 20: %d\n", greater20)
	} else {
		fmt.Printf("В slice нет числа больше 20")
	}
}

func sumSlice(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
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

func findFirstGreaterThan(nums []int, limit int) (int, bool) {
	for _, v := range nums {
		if v > limit {
			return v, true
		}
	}

	return 0, false
}
