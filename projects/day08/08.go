package main

import "fmt"

func main() {
	s := "go-lang-is-fun"
	runes := []rune(s)
	res := make([]rune, 0, len(runes))

	for _, r := range runes {
		if r != '-' {
			res = append(res, r)
		}
	}

	fmt.Println(string(res))
}
