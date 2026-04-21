package main

import "fmt"

func main() {
	s := " g o l a n g "
	runes := []rune(s)
	res := make([]rune, 0, len(runes))

	for _, r := range runes {
		if r != ' ' {
			res = append(res, r)
		}
	}

	s = string(res)

	fmt.Println(s)
	//здесь удобно собирать новую строку потому что строки в go неизменяемые, и чтобы изменить строку, нужно создать новую
}
