package main

import "fmt"

func main() {
	fmt.Println("func main")
	f2()
	defer fmt.Println("main defer")
}

func f2() {
	defer fmt.Println("f2 defer")
	fmt.Println("func f2")
}
