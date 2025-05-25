package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	// need to implement
	errs []error
}

func (e *MultiError) Error() string {
	res := fmt.Sprintf("%d errors occured:\n", len(e.errs))
	for _, err := range e.errs {
		res += fmt.Sprintf("\t* %v", err)
	}
	return res + "\n"
}

func Append(err error, errs ...error) *MultiError {
	var me *MultiError
	if err == nil {
		me = &MultiError{}
	} else {
		me = err.(*MultiError)
	}
	me.errs = append(me.errs, errs...)

	return me
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
