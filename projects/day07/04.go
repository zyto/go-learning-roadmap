package main

import "fmt"

func main() {
	m := map[string]int{"a": 1}

	//читаем несуществующий ключ, выводит 0, без ошибки
	fmt.Println(m["b"])

	//теперь сделаем через ок-идиому, и проверим есть ключ или нет
	val, ok := m["b"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found") //т.к. ключа нет мы можем точно обработать эту ситуацию
		//но так же можем вывести пустое значение:
		fmt.Println(m["b"]) //тут тоже 0 и нет ошибки
	}
}
