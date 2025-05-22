package main

import (
	"errors"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	sb := new(strings.Builder)

	sb.WriteString(fmt.Sprintf("%d errors occured:\n", len(e.errors)))

	for _, err := range e.errors {
		sb.WriteString("\t* ")
		sb.WriteString(err.Error())

	}
	if len(e.errors) > 0 {
		sb.WriteString("\n")
	}

	return sb.String()
}

func Append(err error, errs ...error) *MultiError {

	if err == nil && len(errs) == 0 {
		return nil
	}

	if mErr, ok := err.(*MultiError); ok {
		mErr.errors = append(mErr.errors, errs...)
		return mErr
	}

	multiErr := new(MultiError)
	if err != nil {
		multiErr.errors = append(multiErr.errors, err)
	}
	multiErr.errors = append(multiErr.errors, errs...)

	return multiErr
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
