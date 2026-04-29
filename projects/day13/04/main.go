package main

import "fmt"

func main() {
	value := 100

	p := &value

	fmt.Printf("%T\n", value)
	fmt.Printf("%T\n", p)
	fmt.Printf("%d\n", *p)
}
