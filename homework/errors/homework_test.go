package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MyErr struct {
	msg   string
	cause error
}

func (e MyErr) Error() string {
	if e.cause == nil {
		return e.msg
	}

	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e MyErr) Wrap(msg string) error {
	e.cause = MyErr{
		msg:   e.msg,
		cause: e.cause,
	}
	e.msg = msg

	return e
}

func (e MyErr) Unwrap() error {
	return e.cause
}

func (e MyErr) Is(err error) bool {
	if e == err {
		return e.cause.Error() == err.Error()
	}

	if err == nil {
		return false
	}

	wrappedErr := e.Unwrap()

	for wrappedErr != nil {
		if wrappedErr == err {
			return err.Error() == wrappedErr.Error()
		}

		unWrappableErr, ok := wrappedErr.(interface {
			Unwrap() error
		})
		if !ok {
			wrappedErr = nil
			continue
		}

		wrappedErr = unWrappableErr.Unwrap()
	}

	return false
}

func (e MyErr) As(err any) bool {
	for err != nil {
		if _, ok := err.(*MyErr); ok {
			return true
		}

		unWrappableErr, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			return false
		}

		err = unWrappableErr.Unwrap()
	}

	return false
}

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if len(e.errors) == 0 {
		return ""
	}

	if len(e.errors) == 1 {
		return e.errors[1].Error()
	}

	expectedMsgLen := 0
	errorsCount := 0

	for _, err := range e.errors {
		if err != nil {
			expectedMsgLen += len(err.Error())
			errorsCount++
		}
	}

	if errorsCount == 0 {
		return ""
	}

	template := fmt.Sprintf("%d errors occurred:\n", errorsCount)
	msgPrefix := "\t* "
	endLine := "\n"
	res := strings.Builder{}
	res.Grow(expectedMsgLen + len(template) + len(msgPrefix)*errorsCount + len(endLine))
	res.Write([]byte(template))

	for _, err := range e.errors {
		if err != nil {
			res.Write([]byte(msgPrefix + err.Error()))
		}
	}

	res.Write([]byte(endLine))

	return res.String()
}

func Append(err error, errs ...error) *MultiError {
	var multiErr *MultiError
	if !errors.As(err, &multiErr) {
		multiErr = &MultiError{}
		multiErr.errors = append(multiErr.errors, err)
	}

	for _, err2 := range errs {
		multiErr.errors = append(multiErr.errors, err2)
	}

	return multiErr
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occurred:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)

	originalErr := errors.New("original err")
	myErr := MyErr{"my Err", originalErr}

	err = Append(err, myErr)
	err = Append(err, errors.New("error 4"))

	expectedMessage = "4 errors occurred:\n\t* error 1\t* error 2\t* my Err: original err\t* error 4\n"
	assert.EqualError(t, err, expectedMessage)

	assert.True(t, myErr.Is(originalErr))

	var myErr2 MyErr
	assert.True(t, myErr.As(&myErr2))
}
