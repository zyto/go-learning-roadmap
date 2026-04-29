package main

import "fmt"

func main() {
	m := make(map[string]int)

	addScore(m, "aaa", 5)
	addScore(m, "bbb", 99)

	fmt.Println(m)
}

func addScore(scores map[string]int, name string, score int) {
	scores[name] = score
}
