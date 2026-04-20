package main

import "fmt"

func main() {
	ids := []int{101, 205, 330, 400}
	names := []string{"Alice", "Bob", "Carol"}

	res := make(map[int]string)

	for i, id := range ids {
		res[id] = names[i]
	}

	fmt.Println(res[330])

	name, ok := res[999]
	if !ok {
		fmt.Println("Name not found")
		//мы можем вывести имя для отсутствующего ID, но будет пустая строка:
		fmt.Println(name)
	} else {
		fmt.Println(name)
	}
}
