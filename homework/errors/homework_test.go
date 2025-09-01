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
	Errors []error
}

func (e *MultiError) Error() string {
	if e == nil || len(e.Errors) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d errors occured:\n", len(e.Errors)))

	for _, err := range e.Errors {
		b.WriteString("\t* ")
		b.WriteString(err.Error())
	}
	b.WriteString("\n")

	return b.String()
}

func Append(err error, errs ...error) *MultiError {
	var me *MultiError

	if err == nil {
		me = &MultiError{}
	} else {
		if existing, ok := err.(*MultiError); ok {
			me = existing
		} else {
			me = &MultiError{Errors: []error{err}}
		}
	}

	for _, e := range errs {
		if e != nil {
			me.Errors = append(me.Errors, e)
		}
	}

	if len(me.Errors) == 0 {
		return nil
	}
	return me
}
func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
