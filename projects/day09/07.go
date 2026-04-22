package main

import (
	"fmt"
	"maps"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3}
	b := map[string]int{"one": 1, "two": 2, "three": 3}
	c := map[string]int{"three": 3, "one": 1, "two": 2}

	fmt.Println(maps.Equal(a, b))
	fmt.Println(maps.Equal(b, c))
	//все мапы равны потому что map - это неупорядоченная структура, и Equal сравнивает по ключам и значению, но не по порядку
}
