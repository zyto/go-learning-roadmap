package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := divide(5, 3)
	fmt.Println(res, err)

	res, err = divide(-5, 0)
	fmt.Println(res, err)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
