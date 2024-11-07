package main

import (
	"errors"
	"fmt"
)

func GetRoute(lat, lon float32) (string, error) {
	err := validateCoordinates(lat, lon)
	if err != nil {
		fmt.Println("incorrect coordinates") // again
		return "", err
	}

	return "route", nil
}

func validateCoordinates(lat, lon float32) error {
	if lat > 90. || lat < -90. {
		fmt.Println("incorrect latitude")
		return errors.New("incorrect latitude")
	}
	if lon > 180. || lon < -180. {
		fmt.Println("incorrect longitude")
		return errors.New("incorrect longitude")
	}

	return nil
}
