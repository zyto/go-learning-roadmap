package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{10, 20, 30}
	b := []int{10, 20, 30}
	c := []int{30, 20, 10}

	if slices.Equal(a, b) {
		fmt.Println("a и b равны")
	}

	if slices.Equal(b, c) {
		fmt.Println("b и c равны")
	}

	if slices.Equal(a, c) {
		fmt.Println("a и c равны")
	}

	//срезы сравниваются не только по равенству числе входящих в срез, но и по длине среза и порядку значений
	//поэтому если срезы содержат одинаковые числа в разных позициях - они будут не равны
}
