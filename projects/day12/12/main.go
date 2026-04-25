package main

import (
	"errors"
	"fmt"

	"example.com/day12mini/numbers"
)

func main() {
	res, err := numbers.SumFromStrings([]string{"10", ""})
	if err != nil {
		if errors.Is(err, numbers.ErrEmptyNumber) {
			fmt.Println("пустое число!")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Сумма всех чисел: %d\n", res)
	}
}
