package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, err := parseAge("42")
	fmt.Println(res, err)

	res, err = parseAge("   ")
	fmt.Println(res, err)

	res, err = parseAge("-7")
	fmt.Println(res, err)

	res, err = parseAge("12t")
	fmt.Println(res, err)
}

func parseAge(text string) (int, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, errors.New("age is empty")
	}

	age, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}

	if age < 0 {
		return 0, errors.New("age must be >= 0")
	}

	return age, nil
}
