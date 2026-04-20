package main

import "fmt"

func main() {
	s := []string{"go", "php", "go", "sql", "go", "php"}
	m := make(map[string]bool)
	var res []string

	for _, w := range s {
		if m[w] == false {
			m[w] = true
			res = append(res, w)
		}
	}

	fmt.Println(res)
}
