package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if len(e.errors) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%d errors occured:\n", len(e.errors))
	for _, err := range e.errors {
		fmt.Fprintf(&b, "\t* %s", err.Error())
	}
	b.WriteByte('\n')
	return b.String()
}

func Append(err error, errs ...error) *MultiError {
	var multi *MultiError
	if err != nil {
		if m, ok := err.(*MultiError); ok {
			multi = m
		} else {
			multi = &MultiError{errors: []error{err}}
		}
	} else {
		multi = &MultiError{}
	}

	for _, e := range errs {
		if e != nil {
			multi.errors = append(multi.errors, e)
		}
	}

	if len(multi.errors) == 0 {
		return nil
	}
	return multi
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
