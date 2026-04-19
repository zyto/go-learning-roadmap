package main

import "fmt"

func main() {
	nums := []int{3, 6, 9, 12}

	for i, v := range nums {
		fmt.Printf("Индекс = %d, Значение = %d\n", i, v)
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Printf("Сумма всех элементов = %d\n", sum)
}
