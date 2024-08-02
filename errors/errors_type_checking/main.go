package main

import (
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
	switch err := err.(type) {
	case DatabaseError:
		fmt.Println(err.Error())
	default:
		fmt.Println("unknown error")
	}

}
