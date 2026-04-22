package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i := 42
	s := strconv.Itoa(i)
	fmt.Println(s)

	s = " 105 "
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Преобразование не удалось!")
	} else {
		fmt.Println(i)
	}
}
