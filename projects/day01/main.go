package main

import "fmt"

func main() {
	var name string
	var age int
	var plan int

	fmt.Print("Как вас зовут? ")
	fmt.Scanln(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)

	fmt.Print("Сколько часов в день вы будете уделять Go? ")
	fmt.Scanln(&plan)

	fmt.Printf("Имя: %s\nВозраст: %d\nПлан по Go: %d часов в день\n",
		name,
		age,
		plan)
}
