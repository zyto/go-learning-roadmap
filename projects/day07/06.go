package main

import "fmt"

func main() {
	words := []string{"go", "php", "go", "sql", "go", "php"}
	m := map[string]int{}

	for _, w := range words {
		//даже если в map нет ключа - при получении значения несуществующего ключа будет ноль, для каждого найденного слова мы увеличиваем счётчик на 1
		m[w]++
	}

	fmt.Println(m)
}
