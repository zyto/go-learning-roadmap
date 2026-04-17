package main

import "fmt"

func main() {
	a := 10
	login := "admin"
	password := "1234"

	if login == "admin" && password == "1234" {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Неверные данные")
	}
}
