package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	slice = append(slice, 5)

	//append(slice, 6)

	fmt.Println(slice)
}
