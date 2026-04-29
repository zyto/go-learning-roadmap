package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	b := a

	b[0] = 99

	fmt.Println(a)
	fmt.Println(b)

}
