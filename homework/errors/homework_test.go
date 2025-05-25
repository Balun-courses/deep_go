package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	// if len(e.Errors) == 1 {
	// 	return fmt.Sprintf("1 error occurred:\n\t* %s\n", e.Errors[0])
	// }

	length := len(e.errors)
	points := make([]string, length)
	for i, err := range e.errors {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf("%d errors occured:\n\t%s\n", length, strings.Join(points, "\t"))
}

func Append(err error, errs ...error) *MultiError {
	multiErr, ok := err.(*MultiError)
	if ok {
		multiErr.errors = append(multiErr.errors, errs...)
		return multiErr
	}

	return &MultiError{errors: errs}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
