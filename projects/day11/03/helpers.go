package main

import "strings"

func normalizeWord(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}
