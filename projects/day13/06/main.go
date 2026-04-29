package main

import "fmt"

func main() {
	value := 42
	reset(&value)

	fmt.Printf("%d\n", value)
}

func reset(n *int) {
	*n = 0
}
