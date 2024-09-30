package main

import (
	"errors"
	"fmt"
)

var ErrDatabaseProblem = errors.New("database problem")

func GetDataFromDB() error {
	return fmt.Errorf("failed to get data: %w", ErrDatabaseProblem)
}

func main() {
	err := GetDataFromDB()
	if errors.Is(err, ErrDatabaseProblem) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("unknown error")
	}
}
