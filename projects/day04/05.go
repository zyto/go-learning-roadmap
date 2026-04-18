package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		if i == 4 || i == 8 || i == 12 {
			continue
		}

		fmt.Println(i)
	}
}
