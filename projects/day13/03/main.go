package main

import "fmt"

func main() {
	value := 10

	value = setToHundred(value)

	fmt.Println(value)
}

func setToHundred(n int) int {
	n = 100

	return n
}
