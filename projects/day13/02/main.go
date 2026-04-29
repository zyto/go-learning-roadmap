package main

import "fmt"

func main() {
	value := 10

	setToHundred(value)

	fmt.Println(value)
}

func setToHundred(n int) {
	n = 100
}
