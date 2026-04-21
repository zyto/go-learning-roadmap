package main

import "fmt"

func main() {
	s := "level"
	m := map[rune]int{}

	for _, r := range s {
		m[r]++
	}

	fmt.Println(m['l'])
	fmt.Println(m['e'])
	fmt.Println(m['v'])
}
