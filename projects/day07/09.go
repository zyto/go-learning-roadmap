package main

import "fmt"

func main() {
	departments := []string{"backend", "frontend", "backend", "qa", "frontend"}
	names := []string{"Ira", "Oleg", "Anton", "Lena", "Masha"}

	res := make(map[string][]string)

	for i, department := range departments {
		res[department] = append(res[department], names[i])
	}

	fmt.Println(res)

	//здесь удобно использовать именно map[K][]V потому что в отделе может быть не один человек а несколько (или 0), поэтому значение map у нас не string а срез из строк
}
