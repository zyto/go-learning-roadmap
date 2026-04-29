package main

import "fmt"

func main() {
	printValue(nil)
	value := 15
	printValue(&value)

}

func printValue(p *int) {
	if p == nil {
		fmt.Println("no value")
		return
	}

	fmt.Printf("%d\n", *p)
}
