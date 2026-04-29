package main

import "fmt"

func main() {
	a := 10
	b := a

	b = 100
	fmt.Println(a)
	fmt.Println(b)
}
