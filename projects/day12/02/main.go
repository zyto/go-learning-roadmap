package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := strconv.Atoi("42")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
