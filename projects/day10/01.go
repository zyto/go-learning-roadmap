package main

import "fmt"

func main() {
	fmt.Println(sumTwo(10, 20))
	fmt.Println(sumTwo(-5, 3))
}

func sumTwo(a, b int) int {
	return a + b
}
