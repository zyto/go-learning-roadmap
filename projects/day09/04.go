package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{50, 10, 40, 20, 30}

	slices.Sort(nums)

	fmt.Println(nums) //здесь мы увидим что изначальный срез отсортирован

	//если мы хотим сохранить изначальный срез то мы должны создать копию и её уже сортировать
	nums = []int{50, 10, 40, 20, 30}
	dst := slices.Clone(nums)

	slices.Sort(dst)
	fmt.Println(nums)
	fmt.Println(dst)
}
