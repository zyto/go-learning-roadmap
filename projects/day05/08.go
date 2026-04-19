package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 10)
	slice2 := make([]int, 0)

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000)
	}

	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			slice2 = append(slice2, slice[i])
		}
	}

	fmt.Println(slice2)
}
