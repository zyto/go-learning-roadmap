package main

import "fmt"

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	case 3:
		fmt.Println("Три")
	default:
		fmt.Println("Другое")
	}
}
