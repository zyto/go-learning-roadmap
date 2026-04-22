package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := "report-final.json"
	route := "/api/v1/users"

	sub := "final"
	fmt.Println(strings.Contains(filename, sub))

	suffix := ".json"
	fmt.Println(strings.HasSuffix(filename, suffix))

	prefix := "/api/"
	fmt.Println(strings.HasPrefix(route, prefix))
}
