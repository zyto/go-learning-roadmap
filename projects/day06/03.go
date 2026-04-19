package main

import "fmt"

func main() {
	var a []int
	b := []int{}

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(cap(a))
	fmt.Println(cap(b))

	//оба среза имеют нулевую длину и нулевую вместимость
	//однако только первый будет равен nil, второй будет просто пустой срез, но не равен nil
}
