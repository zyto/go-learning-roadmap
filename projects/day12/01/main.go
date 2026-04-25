package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := strconv.Atoi("42")

	fmt.Println(err == nil)
	fmt.Println(res)
}
