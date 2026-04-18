package main

import "fmt"

func main() {
	temp := 15
	switch {
	case temp < 0:
		fmt.Println("Холодно")
	case temp < 20:
		fmt.Println("Прохладно")
	default:
		fmt.Println("Тепло")
	}
}
