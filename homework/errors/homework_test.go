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
	if e == nil {
		return ""
	}

	messages := make([]string, 0, len(e.errs))
	for i := range e.errs {
		if e.errs[i] == nil {
			continue
		}
		messages = append(messages, "\t* "+e.errs[i].Error())
	}

	if len(e.errs) == 0 || len(messages) == 0 {
		return ""
	}

	return fmt.Sprintf("%d errors occured:\n", len(messages)) + strings.Join(messages, "") + "\n"
}

func Append(err error, errs ...error) *MultiError {
	if err == nil {
		return &MultiError{errs: errs}
	}

	if mError, ok := err.(*MultiError); ok {
		mError.errs = append(mError.errs, errs...)
		return mError
	}

	return &MultiError{errs: append([]error{err}, errs...)}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
