package numbers

import (
	"errors"
	"strconv"
	"strings"
)

var ErrEmptyNumber = errors.New("empty number")

func ParsePositive(text string) (int, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return 0, ErrEmptyNumber
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}

	if val <= 0 {
		return 0, ErrEmptyNumber
	}

	return val, nil
}
