package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errList []error
}

func (e *MultiError) Error() string {
	bld := strings.Builder{}
	bld.WriteString("2 errors occured:\n")
	for _, err := range e.errList {
		if err != nil {
			bld.WriteString("\t* ")
			bld.WriteString(err.Error())
		}
	}
	bld.WriteString("\n")
	return bld.String()
}

func Append(err error, errs ...error) *MultiError {
	var multiErr *MultiError

	if errors.As(err, &multiErr) {
		multiErr, _ = err.(*MultiError)
	} else {
		multiErr = &MultiError{errList: make([]error, 0)}
		multiErr.errList = append(multiErr.errList, err)
	}

	for _, e := range errs {
		multiErr.errList = append(multiErr.errList, e)
	}

	return multiErr
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
