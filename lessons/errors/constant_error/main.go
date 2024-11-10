package main

import (
	"errors"
	"fmt"
	"io"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const ErrEOF = Error("EOF")

func main() {
	var err error = ErrEOF
	err = io.EOF
	_ = err

	// ErrEOF = Error("new error") -> coplilation error

	anotherErrEOF1 := Error("EOF")
	fmt.Println(`anotherErrEOF1 = Error("EOF"):`, anotherErrEOF1 == ErrEOF)

	anotherErrEOF2 := errors.New("EOF")
	fmt.Println("anotherErrEOF2 = io.EOF:", anotherErrEOF2 == io.EOF)
}
