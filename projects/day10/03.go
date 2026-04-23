package main

import "fmt"

func main() {

	i, ok := findIndex([]int{1, 2, 3}, 2)

	if !ok {
		fmt.Println("Поиск невозможен")
	} else if i >= 0 {
		fmt.Printf("Найден индекс: %d\n", i)
	} else {
		fmt.Println("Значение не найдено")
	}
}

func findIndex(nums []int, target int) (int, bool) {
	for i, n := range nums {
		if n == target {
			return i, true
		}
	}

	return -1, false

	//false полезен чтобы показать что поиск невозможен, например потому что срез пустой - и поиск нет смысла выполнять
	//-1 говорит что срез не пустой, но переданного значения не содержит
	//ну а какое-то значение от нуля и больше - это индекс найденного значения
}
