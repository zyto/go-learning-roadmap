package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30

	m["key1"] = 999
	fmt.Println(m)

	//при создании через литерал мы можем сразу задать ключи и значения, а make создаёт пустой map
}
