package main

import (
	"fmt"
)

func main() {
	age := 16

	if age < 14 {
		fmt.Println("Ребёнок")
	} else if age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
}
