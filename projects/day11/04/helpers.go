package main

import "strings"

func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}

func normalizeAll(words []string) []string {
	dst := make([]string, 0, len(words))

	for _, w := range words {
		dst = append(dst, normalizeWord(w))
	}

	return dst
}
