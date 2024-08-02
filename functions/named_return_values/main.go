package main

import (
	"context"
	"errors"
)

type mobileNavigator struct{}

func (n mobileNavigator) GetLocation(ctx context.Context, address string) (lat, lon float32, err error) {
	if address == "" {
		return 0., 0., errors.New("incorrect address")
	}

	if ctx.Err() != nil {
		return 0., 0., err // nil error
	}

	// calculating coordinates...
	return lat, lon, err
}
