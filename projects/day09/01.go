package main

import (
	"fmt"
	"slices"
)

func main() {
	langs := []string{"go", "php", "sql", "redis"}

	search := "go"
	if slices.Contains(langs, search) {
		fmt.Println("В langs есть " + search)
	} else {
		fmt.Println("В langs нет " + search)
	}

	search = "java"
	if slices.Contains(langs, search) {
		fmt.Println("В langs есть " + search)
	} else {
		fmt.Println("В langs нет " + search)
	}

	search = "sql"
	fmt.Println(slices.Index(langs, search))

	search = "docker"
	fmt.Println(slices.Index(langs, search))
	//индекс -1 потому что первый индекс = 0, и было бы не понятно ноль - это отсутствующее значение или первый индекс
	//а -1 индекса не может быть, поэтому такое значение точно даёт понять что такого значения не существует
}
