package main

import (
	"fmt"

	"example.com/day11task09/words"
)

func main() {
	s := []string{" Go ", " SQL ", "go", " php "}

	n := words.NormalizeAll(s)
	fmt.Println(n)

	fmt.Println(words.CountUnique(n))
}
