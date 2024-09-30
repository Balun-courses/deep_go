package main

import (
	"errors"
	"fmt"
)

func createErrV1() error {
	return errors.New("simple error")
}

func createErrV2() error {
	return fmt.Errorf("error with argument %d", 1000)
}

func main() {
	fmt.Println(createErrV1().Error())
	fmt.Println(createErrV2().Error())
}
