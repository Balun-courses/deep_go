package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var s string
	for _, err := range e.errs {
		s += fmt.Sprintf("\t* %v", err)
	}

	return fmt.Sprintf("%d errors occured:\n%s\n", len(e.errs), s)
}

func (e *MultiError) Unwrap() []error {
	return e.errs
}

func Append(err error, errs ...error) *MultiError {
	var result *MultiError

	if err != nil {
		if multi, ok := err.(*MultiError); ok {
			result = multi
		} else {
			result = &MultiError{errs: []error{err}}
		}
	} else {
		result = &MultiError{}
	}

	for _, e := range errs {
		if e != nil {
			result.errs = append(result.errs, e)
		}
	}

	return result
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiErrorUnwrap(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	err := Append(nil, err1, err2)

	assert.True(t, errors.Is(err, err1))
	assert.True(t, errors.Is(err, err2))

	err3 := errors.New("error 3")
	assert.False(t, errors.Is(err, err3))
}
