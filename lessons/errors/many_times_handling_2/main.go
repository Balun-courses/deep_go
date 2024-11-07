package main

import (
	"errors"
	"fmt"
)

func GetRoute(lat, lon float32) (string, error) {
	err := validateCoordinates(lat, lon)
	if err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}

	return "route", nil
}

func validateCoordinates(lat, lon float32) error {
	if lat > 90. || lat < -90. {
		return errors.New("incorrect latitude")
	}
	if lon > 180. || lon < -180. {
		return errors.New("incorrect longitude")
	}

	return nil
}
