package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[2:4]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
