package main

import (
	"fmt"
	"strings"

	"example.com/day11task10/words"
)

func main() {
	s := " go, php ,sql, go "

	//можно добавить strings.Split в CountUnique - но тогда эта функция будет делать слишком много разных действий
	//поэтому, как мне кажется, лучше сделать как я сделал
	//и отдельную функцию для одной строки нет смысла делать
	fmt.Println(words.CountUnique(strings.Split(s, ",")))
}
