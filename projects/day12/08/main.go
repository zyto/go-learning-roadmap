package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	_ = process("")
	_ = process("10")
}

func process(text string) error {
	fmt.Println("begin")
	defer fmt.Println("cleanup")
	if strings.TrimSpace(text) == "" {
		return errors.New("empty text")
	}
	fmt.Println("processing...")
	return nil
}
