package main

import (
	"context"
)

type mobileNavigator struct{}

func (n mobileNavigator) GetLocation(ctx context.Context, address string) (lat, lon float32, err error) {
	if address == "" {
		return
	}

	if ctx.Err() != nil {
		return
	}

	// calculating coordinates...
	return lat, lon, err
}
