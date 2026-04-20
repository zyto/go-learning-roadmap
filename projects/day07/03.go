package main

import "fmt"

func main() {
	var scores map[string]int

	fmt.Println(scores == nil)
	fmt.Println(scores["math"])

	//выведет ошибку (panic), поэтому закомментировал:
	//scores["math"] = 100

	scores = make(map[string]int)
	scores["math"] = 100 //тут ошибки не будет, этим и отличается nil-map от пустого
}
