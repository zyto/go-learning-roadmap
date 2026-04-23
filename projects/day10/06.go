package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("\"%s\"\n", normalizeWord(" Go "))
	fmt.Printf("\"%s\"\n", normalizeWord(" SQL"))
	fmt.Printf("\"%s\"\n", normalizeWord("php  "))
	//я специально вывел кавычки, чтобы убедиться что лишние пробелы убраны
}

func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}
