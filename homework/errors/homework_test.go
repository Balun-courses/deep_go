package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// MultiError объединяет несколько ошибок в одну.
type MultiError struct {
	Errors []error
}

// Error возвращает удобочитаемое представление всех накопленных ошибок.
func (e *MultiError) Error() string {
	if len(e.Errors) == 0 {
		return "0 errors occured:\n"
	}

	if len(e.Errors) == 1 {
		return fmt.Sprintf("1 error occured:\n\t* %s\n", e.Errors[0])
	}

	msg := fmt.Sprintf("%d errors occured:\n", len(e.Errors))
	for _, err := range e.Errors {
		msg += fmt.Sprintf("\t* %s", err)
	}
	return msg + "\n"
}

// Append добавляет ошибки к err. Если err — это *MultiError, ошибки добавляются к нему,
//в противном случае создается новый *MultiError. Ошибки со значением nil пропускаются.
func Append(err error, errs ...error) *MultiError {
	var result *MultiError

	switch typed := err.(type) {
	case *MultiError:
		if typed != nil {
			result = typed
		}
	case nil:
		// nothing to do
	default:
		result = &MultiError{Errors: []error{err}}
	}

	appendOne := func(e error) {
		if e == nil {
			return
		}
		if result == nil {
			result = &MultiError{}
		}
		if nested, ok := e.(*MultiError); ok && nested != nil {
			result.Errors = append(result.Errors, nested.Errors...)
			return
		}
		result.Errors = append(result.Errors, e)
	}

	for _, e := range errs {
		appendOne(e)
	}

	return result
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiErrorSingle(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))

	expectedMessage := "1 error occured:\n\t* error 1\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiErrorEmpty(t *testing.T) {
	var err error
	err = Append(err, nil)

	assert.Nil(t, err)
}

func TestAppendToPlainError(t *testing.T) {
	err := Append(errors.New("error 1"), errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestAppendNilErrorsSkipped(t *testing.T) {
	err := Append(nil, nil, errors.New("error 1"), nil)

	assert.Len(t, err.Errors, 1)
	assert.EqualError(t, err, "1 error occured:\n\t* error 1\n")
}

func TestAppendMultiErrorFlattened(t *testing.T) {
	first := Append(nil, errors.New("error 1"))
	second := Append(nil, errors.New("error 2"), errors.New("error 3"))

	err := Append(first, second)

	assert.Len(t, err.Errors, 3)
	assert.EqualError(t, err, "3 errors occured:\n\t* error 1\t* error 2\t* error 3\n")
}

func TestAppendManyAtOnce(t *testing.T) {
	err := Append(nil, errors.New("a"), errors.New("b"), errors.New("c"))

	assert.Len(t, err.Errors, 3)
}

func TestMultiErrorImplementsError(t *testing.T) {
	var err error = &MultiError{Errors: []error{errors.New("x")}}

	assert.NotPanics(t, func() { _ = err.Error() })
}
