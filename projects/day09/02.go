package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c := append(slices.Clone(a), b...)

	fmt.Println(a, b, c)
}
