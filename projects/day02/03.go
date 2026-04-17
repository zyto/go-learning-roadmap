package main

import (
	"fmt"
)

func main() {
	age := 20
	hasTicket := true

	if age >= 18 && hasTicket {
		fmt.Println("Проход разрешён")
	} else {
		fmt.Println("Проход запрещён")
	}
}
