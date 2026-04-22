package main

import (
	"fmt"
	"maps"
)

func main() {
	defaults := map[string]string{
		"host": "localhost",
		"port": "8080",
		"env":  "dev",
	}

	user := map[string]string{
		"port": "9090",
	}

	dst := maps.Clone(defaults)

	maps.Copy(dst, user)

	fmt.Println(defaults)
	fmt.Println(dst)
}
