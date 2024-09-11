package main

import "io"

type T struct {
	//...
}

// Interface guard
var _ io.ReadWriter = (*T)(nil)

func (t *T) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (t *T) Write(p []byte) (n int, err error) {
	return 0, nil
}
