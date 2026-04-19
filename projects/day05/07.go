package main

import "fmt"

func main() {
	nums1 := [3]int{1, 2, 3}
	nums2 := nums1

	nums2[1] = 99

	fmt.Println(nums1)
	fmt.Println(nums2)

	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2[1] = 99

	fmt.Println(slice1)
	fmt.Println(slice2)

}
