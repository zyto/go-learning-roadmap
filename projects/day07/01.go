package main

import "fmt"

func main() {
	goods := map[string]int{
		"computer": 1500,
		"table":    10,
	}

	goods["cellphone"] = 150

	fmt.Println(goods)
	fmt.Println(goods["cellphone"])
}
