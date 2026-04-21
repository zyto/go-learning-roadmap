package main

import "fmt"

func main() {
	s := "кот"

	for i, r := range s {
		fmt.Println(i)
		fmt.Println(r)
	}

	//индексы идут как 0, 2, 4 потому что они выводят индекс байта а не символа
}
