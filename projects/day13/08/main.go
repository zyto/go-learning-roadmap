package main

import "fmt"

func main() {

	a := 10
	b := 10

	p1 := &a
	p2 := &a
	p3 := &b

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)
}
