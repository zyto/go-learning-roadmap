package main

import "fmt"

func main() {
	process()
}

func process() {
	fmt.Println("start")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("finish body")
}
