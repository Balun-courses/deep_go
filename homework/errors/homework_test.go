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
	errs []error
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}
	var result strings.Builder
	count := len(e.errs)
	if count == 1 {
		return e.errs[0].Error()
	}
	fmt.Fprintf(&result, "%d errors occured:\n", count)
	result.WriteString("\t* ")
	result.WriteString(e.errs[0].Error())
	for _, err := range e.errs[1:] {
		result.WriteString("\t* ")
		result.WriteString(err.Error())
	}
	result.WriteString("\n")
	return result.String()
}

func Append(err error, errs ...error) *MultiError {
	var result []error

	if err != nil {
		if multi, ok := err.(*MultiError); ok {
			result = append(result, multi.errs...)
		} else {
			result = append(result, err)
		}
	}

	for _, err := range errs {
		if err != nil {
			if multi, ok := err.(*MultiError); ok {
				result = append(result, multi.errs...)
			} else {
				result = append(result, err)
			}
		}
	}

	if len(result) == 0 {
		return nil
	}

	return &MultiError{errs: result}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
