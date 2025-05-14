package main

import (
	"errors"
	"strconv"
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
	res := strings.Builder{}
	res.WriteString(strconv.Itoa(len(e.errs)))
	res.WriteString(" errors occured:\n")

	for _, err := range e.errs {
		res.WriteString("\t* ")
		res.WriteString(err.Error())
	}

	res.WriteString("\n")

	return res.String()
}

func Append(err error, errs ...error) *MultiError {
	if multiErr, ok := err.(*MultiError); ok {
		multiErr.errs = append(multiErr.errs, errs...)
		return multiErr
	}
	return &MultiError{
		errs: errs,
	}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
