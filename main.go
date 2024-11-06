package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		return
	}

	fileName := os.Args[1]

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("File '%s' already exists.\n", fileName)
		} else {
			fmt.Printf("Error creating file: %v\n", err)
		}
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' created successfully.\n", fileName)
}
