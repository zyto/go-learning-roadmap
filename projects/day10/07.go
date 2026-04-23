package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{" Go ", " SQL", "php  "}

	fmt.Println(normalizeWords(words))
}

func normalizeWords(words []string) []string {
	dst := make([]string, 0, len(words))
	for _, word := range words {
		dst = append(dst, normalizeWord(word))
	}

	return dst
}

// т.к. я делаю все задачи - каждую в своём файле - то я скопировал эту функцию сюда, в файл с новой задачей
// возможно мне надо было как-то подключить функцию из прошлого урока, но я не знаю как, этого не было в уроках
// codex, подумай и предложи как лучше в дальнейшем организовывать выполнение задач, чтобы новая задача не затирала старую, и я мог отдать на ревью все задачи
// я мог бы весь код этого файла написать в файле прошлой задачи в 06.go, тогда бы я выполнил условие "не дублировать логику из предыдущей задачи, а переиспользовать helper function"
// но при этом в 6 задаче появился бы лишний код
func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}
