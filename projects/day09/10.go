package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := "   Admin@Example.COM   "
	dst := strings.ToLower(strings.TrimSpace(raw))

	fmt.Println(dst == raw)
}
