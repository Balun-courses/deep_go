package main

import (
	"errors"
	"fmt"
)

type DatabaseError struct{}

func (d DatabaseError) Error() string {
	return "database error"
}

func GetDataFromDB() error {
	return fmt.Errorf("failed to get data: %w", DatabaseError{})
}

func main() {
	err := GetDataFromDB()
	if errors.As(err, &DatabaseError{}) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("unknown error")
	}
}
