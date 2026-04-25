package main

import (
	"fmt"
	"os"
)

func main() {
	makeTempFile()
}

func makeTempFile() {
	file, err := os.Create("temp.txt")

	if err != nil {
		fmt.Println("create error:", err)
		return
	}

	defer os.Remove("temp.txt")
	defer file.Close()

	fmt.Println("file was created")
}
