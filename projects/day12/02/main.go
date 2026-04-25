package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := strconv.Atoi("42x")

	fmt.Println(err == nil)
	fmt.Println(err)
	fmt.Println(res)
	//числом нельзя пользоваться потому что err != nill, т.е. функция возвратила ошибку, это значит что число некорректное, скорее всего там пустое значение для типа, т.е. для int это 0

}
