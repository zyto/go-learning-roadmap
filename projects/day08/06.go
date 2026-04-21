package main

import "fmt"

func main() {
	s := "кот"

	runes := []rune(s)

	runes[0] = 'р'

	s = string(runes)
	fmt.Println(s)
}
