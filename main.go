package main

import (
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args) > 1) {
		fmt.Println("Please enter a filename")
		return
	}

	filename := os.Args[1]

	_, err := os.Create(filename)

	if err != nil {
		fmt.Println("unable to make the file", err.Error())
	}
}
