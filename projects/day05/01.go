package main

import "fmt"

func main() {
	var nums = [5]int{1, 3, 5, 7, 9}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
