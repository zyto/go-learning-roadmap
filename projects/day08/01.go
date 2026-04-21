package main

import "fmt"

func main() {
	s := "golang"

	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	//думаю: тут всё выглядит “просто и логично” потому что строка в ASCII, для UTF-8 было бы по-другому
}
