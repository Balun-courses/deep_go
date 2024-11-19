package main

import (
	"fmt"
	"os"
)

func ReadFile(filename string) error {
	var err *os.PathError

	if filename == "" {
		return err
	}

	// reading...
	return err
}

func main() {
	err := ReadFile("text.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("nil")
	}

	/*
		fmt.Println("value of err: ", err)
		fmt.Printf("type of err: %T\n", err)
		fmt.Println("(err == nil): ", err == nil)
	*/
}
