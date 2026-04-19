package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, -10}
	var result []int

	for _, num := range nums {
		if num >= 0 && num%2 == 0 {
			result = append(result, num)
		}
	}

	fmt.Println(result)
}
