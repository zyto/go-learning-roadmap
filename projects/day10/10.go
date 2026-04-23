package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{
		" www.ya.ru",
		"     google.Com",
		"wWw.YaHoo.com     ",
	}

	//пытаемся найти без нормализации
	fmt.Println(findFirstWithPrefix(words, "www."))

	//теперь запустим функцию с нормализацией
	fmt.Println(findFirstWithPrefixNormalized(words, "wWw."))
}

func findFirstWithPrefixNormalized(words []string, prefix string) (string, bool) {
	words = normalizeWords(words)
	prefix = normalizeWord(prefix)

	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			return word, true
		}
	}

	return "", false
}

func findFirstWithPrefix(words []string, prefix string) (string, bool) {
	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			return word, true
		}
	}

	return "", false
}

func normalizeWords(words []string) []string {
	dst := make([]string, 0, len(words))
	for _, word := range words {
		dst = append(dst, normalizeWord(word))
	}

	return dst
}

// опять-таки здесь я опять скопировал функции из прошлых уроков т.к. не знаю как их использовать повторно, возможно это тема следующих уроков?
func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}
