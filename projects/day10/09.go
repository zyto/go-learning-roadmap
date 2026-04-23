package main

import "fmt"

func main() {
	fmt.Println(maxAll(1, -5, 99))
}

func maxAll(nums ...int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	return m, true
	//здесь недостаточно просто вернуть 0, т.к. может быть срез из одного значения - 0, он и будет максимальным числом
	//и если на пустом срезе мы вернём 0 - это по сути будет тот же результат, и мы не сможем отличить ошибку от описанной ситуации
}
