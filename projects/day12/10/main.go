package main

import "fmt"

func main() {
	value := 10
	defer fmt.Println("deferred: ", value)
	value = 20
	fmt.Println("value: ", value)

	//объяснение: в defer аргументы вычисляются сразу, на момент объявления defer value был 10, он и вывелся в результате
}
