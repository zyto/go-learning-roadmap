package main

import "fmt"

func main() {
	s := "мир"

	bytes := []byte(s)
	runes := []rune(s)

	fmt.Println(bytes)
	fmt.Println(len(bytes))
	fmt.Println(runes)
	fmt.Println(len(runes))

	//bytes важен когда нужно получить чистые байты, а runes нужен когда нужно получить UTF-символы или эмоджи
}
