package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v homework_test.go

type MultiError struct {
	// need to implement
	errors []error
}

func Is(err error, target error) bool {
	val, ok := err.(*MultiError)
	if !ok {
		return false
	}

	for i := range val.errors {
		if reflect.TypeOf(val.errors[i]) != reflect.TypeOf(target) {
			continue
		}

		if reflect.DeepEqual(target, val.errors[i]) {
			return true
		}
	}

	return false
}

func As(err error, target any) bool {
	val, ok := err.(*MultiError)
	if !ok {
		return false
	}

	if target == nil {
		return false
	}

	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return false
	}

	e := t.Elem()
	for i := range val.errors {
		if reflect.TypeOf(val.errors[i]).AssignableTo(e) {
			reflect.ValueOf(target).Elem().Set(reflect.ValueOf(val.errors[i]))
			return true
		}
	}

	return false
}

func Unwrap(err error) error {
	val, ok := err.(*MultiError)
	if !ok {
		return err
	}

	if len(val.errors) <= 1 {
		return nil
	}

	return &MultiError{
		errors: append([]error{}, val.errors[1:]...),
	}
}

func (e *MultiError) Error() string {
	// need to implement
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%d errors occured:\n", len(e.errors)))
	for i := range e.errors {
		b.WriteString("\t* ")
		b.WriteString(e.errors[i].Error())
	}
	b.WriteByte('\n')
	return b.String()
}

func Append(err error, errs ...error) *MultiError {
	// need to implement
	switch val := err.(type) {
	case *MultiError:
		val.errors = append(val.errors, errs...)
		return val
	default:
		var errors []error
		if err != nil {
			errors = append(errors, err)
		}
		return &MultiError{
			errors: append(errors, errs...),
		}
	}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestIs(t *testing.T) {
	e1 := errors.New("foo")
	e2 := errors.New("bar")
	me := Append(nil, e1, e2)

	require.True(t, Is(me, e1), "`errors.Is` должен находить e1")
	require.True(t, Is(me, e2), "`errors.Is` должен находить e2")
	require.False(t, Is(me, errors.New("baz")))
}

type typedErr struct{}

func (typedErr) Error() string { return "typed" }

func TestAs(t *testing.T) {
	var tgt typedErr
	var err error
	me := Append(nil, typedErr{})

	require.True(t, As(me, &tgt), "`As` должен сделать приведение")
	require.False(t, As(me, err), "`As` должен сделать приведение")
}

func TestUnwrap(t *testing.T) {
	e1 := errors.New("first")
	e2 := errors.New("second")
	e3 := errors.New("third")

	me := Append(nil, e1, e2, e3)

	// Первый unwrap отдаёт оставшиеся (e2,e3)
	next := Unwrap(me)
	require.True(t, Is(next, e2))
	require.False(t, Is(next, e1))

	// Ещё раз unwrap → (e3)
	last := Unwrap(next)
	require.True(t, Is(last, e3))

	// Последний unwrap → nil
	require.Nil(t, Unwrap(last))
}
