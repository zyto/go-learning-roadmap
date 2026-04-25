package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(loadPort("10"))
	fmt.Println(loadPort("a101"))
}

func parsePort(text string) (int, error) {
	port, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse port: %w", err)
	}

	return port, nil
}

func loadPort(text string) (int, error) {
	port, err := parsePort(text)
	if err != nil {
		return 0, fmt.Errorf("load port: %w", err)
	}

	return port, nil
}
