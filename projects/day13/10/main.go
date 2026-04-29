package main

import "fmt"

func main() {
	value := [3]int{1, 2, 3}
	changeFirst(value)

	fmt.Println(value)
}

func changeFirst(nums [3]int) {
	nums[0] = 99
}
