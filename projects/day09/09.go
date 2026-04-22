package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "go,php,sql"
	fmt.Println(strings.Split(str, ","))

	str = "  go   php   sql  "
	fmt.Println(strings.Join(strings.Fields(str), " | "))

	//Split лучше когда есть чёткий разделитель
	//Fields лучше когда разделитель - это один или несколько пробелов и нужно убрать лишние пробелы
}
