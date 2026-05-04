package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1

	resetWrong(m)
	fmt.Println(m)

	m = resetRight(m)
	fmt.Println(m)
}

func resetWrong(scores map[string]int) {
	scores = map[string]int{}
}

func resetRight(scores map[string]int) map[string]int {
	return map[string]int{}
}
