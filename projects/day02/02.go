package main

import (
	"fmt"
)

func main() {
	num := -5

	if num > 0 {
		fmt.Println("Положительное")
	} else if num < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}

	if num%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

}
