package main

import "fmt"

func main() {
	value := 10
	p := &value

	*p = 50

	fmt.Println(value)
}
