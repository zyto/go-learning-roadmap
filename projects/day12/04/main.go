package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	words := []string{"go", "php", "sql"}

	i, ok := findWord(words, "php")

	fmt.Println(i, ok)

	limits := []string{"10", "0", "abc"}

	for _, l := range limits {
		limit, err := parseLimit(l)

		fmt.Println(limit, err)
	}

}

func findWord(words []string, target string) (int, bool) {
	for i, word := range words {
		if word == target {
			return i, true
		}
	}

	return -1, false
}

func parseLimit(text string) (int, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, errors.New("empty text")
	}

	val, err := strconv.Atoi(text)

	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}

	if val <= 0 {
		return 0, errors.New("limit must be positive")
	}

	return val, nil
}
