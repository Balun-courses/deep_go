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
	if e == nil || len(e.errors) == 0 {
		return ""
	}

	if len(e.errors) == 1 {
		return e.errors[0].Error()
	}

	var sb strings.Builder
	_, _ = fmt.Fprintf(&sb, "%d errors occured:\n", len(e.errors))
	for _, err := range e.errors {
		_, _ = fmt.Fprintf(&sb, "\t* %v", err)
	}
	sb.WriteString("\n")

	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	if len(errs) == 0 {
		if err == nil {
			return nil
		}
		if me, ok := err.(*MultiError); ok {
			return me
		}
		return &MultiError{errors: []error{err}}
	}

	// Создаем или расширяем MultiError
	var multiErr *MultiError
	if err == nil {
		multiErr = &MultiError{}
	} else {
		switch e := err.(type) {
		case *MultiError:
			multiErr = e
		default:
			multiErr = &MultiError{errors: []error{err}}
		}
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
