package main

import "fmt"

func main() {
	s := "привет"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))

	//len(s) выводит число байт а len([]rune(s)) выводит количество UTF-символов
}
