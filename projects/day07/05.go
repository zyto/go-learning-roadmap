package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
	delete(m, "z") //ошибки не будет, но ничего не сделается, т.к. ключа нет
	fmt.Println(m)
}
