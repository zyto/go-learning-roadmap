package main

import (
	"fmt"
	"slices"
)

func main() {
	scores := map[string]int{
		"bob":    80,
		"alice":  95,
		"claire": 88,
	}

	slice := make([]string, 0, len(scores))

	for k := range scores {
		slice = append(slice, k)
	}

	slices.Sort(slice)

	for _, v := range slice {
		fmt.Printf("%s -> %d\n", v, scores[v])
	}
}
