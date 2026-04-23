package main

import (
	"12/words"
	"fmt"
)

func main() {
	s := []string{" Go ", " SQL ", "go", " php "}

	n := words.NormalizeAll(s)

	fmt.Println(n)
	fmt.Println(words.CountUnique(n))
}
