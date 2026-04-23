package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := " go, php ,sql, go,  php , go "

	//здесь сразу в одном вызове разбиваем строку на slice строк и передаём получившийся срез в функцию построения частотного словаря
	fmt.Println(buildFrequency(splitCSV(raw)))
}

func splitCSV(raw string) []string {
	return strings.Split(raw, ",")
}

// здесь я опять скопировал функции из прошлых уроков т.к. не знаю как их использовать повторно, возможно это тема следующих уроков?
func normalizeWords(words []string) []string {
	dst := make([]string, 0, len(words))
	for _, word := range words {
		dst = append(dst, normalizeWord(word))
	}

	return dst
}

func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}

func filterNonEmpty(words []string) []string {
	dst := make([]string, 0, len(words))
	for _, word := range words {
		//здесь я не проверяю что строка содержит только пробелы, потому что в другой функции я нормализую строку и убираю лишние пробелы
		//можно и здесь проверить, но это будет дублирование кода, а задача стоит именно: разбиение решения на шаги и переиспользование функций
		if word != "" {
			dst = append(dst, word)
		}
	}

	return dst
}

func buildFrequency(words []string) map[string]int {
	dst := make(map[string]int)
	words = normalizeWords(words)
	words = filterNonEmpty(words)
	for _, word := range words {
		dst[word]++
	}

	return dst
}
