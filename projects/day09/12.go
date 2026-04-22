package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	raw := "  go,PHP, go , sql,php,  redis "
	dst := make(map[string]int)

	slice := strings.Split(raw, ",")

	for _, v := range slice {
		dst[strings.ToLower(strings.TrimSpace(v))]++
	}

	slice_sort := make([]string, 0, len(dst))
	for k := range dst {
		slice_sort = append(slice_sort, k)
	}

	slices.Sort(slice_sort)

	for _, v := range slice_sort {
		fmt.Printf("%s: %d\n", v, dst[v])
	}
}
