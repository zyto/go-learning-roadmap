package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyName = errors.New("empty name")

func main() {
	_, err := loadDisplayName("")

	if errors.Is(err, ErrEmptyName) {
		fmt.Println("Ошибка является ошибкой типа", ErrEmptyName)
	}
}

func normalizeName(text string) (string, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return "", ErrEmptyName
	}

	return text, nil
}

func loadDisplayName(text string) (string, error) {
	normalizedName, err := normalizeName(text)

	if err != nil {
		return "", fmt.Errorf("error normalizing name: %w", err)
	}

	return normalizedName, nil
}
