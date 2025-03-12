package main

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (err error) {
	var internalErr *os.PathError

	if filename == "" {
		return internalErr
	}

	// reading...
	return nil
}

func main() {
	err := ReadFile("")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("nil")
	}

	fmt.Println("value of err: ", err)
	fmt.Printf("type of err: %T\n", err)
	fmt.Println("(err == nil): ", err == nil)

}
