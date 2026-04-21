package main

import "fmt"

func main() {
	s := "привет"
	runes := []rune(s)
	res := make([]rune, 0, len(s))

	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}

	s = string(res)

	fmt.Println(s)
}
