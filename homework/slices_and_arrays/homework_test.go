package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func Zero[T any]() (ret T) {
	return
}

type CircularQueue[T any] struct {
	values []T
	front  int
	len    int
}

func NewCircularQueue[T any](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	size := cap(q.values)

	if q.len == size {
		return false
	}

	idx := (q.front + q.len) % size
	q.values[idx] = value
	q.len++

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.len == 0 {
		return false
	}

	size := cap(q.values)

	q.front = (q.front + 1) % size
	q.len--

	return true
}

func (q *CircularQueue[T]) Front() (T, bool) {
	if q.Empty() {
		return Zero[T](), false
	}

	return q.values[q.front], true
}

func (q *CircularQueue[T]) Back() (T, bool) {
	if q.Empty() {
		return Zero[T](), false
	}

	idx := (q.front + q.len - 1) % cap(q.values)

	return q.values[idx], true
}

func (q *CircularQueue[T]) Empty() bool {
	return q.len == 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.len == cap(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	// helper to assert result of funciton with 2 return values
	assertHelper := func(first int, second bool, fn func() (int, bool)) {
		result, ok := fn()
		assert.Equal(t, first, result)
		assert.Equal(t, second, ok)
	}

	assertHelper(0, false, queue.Front)
	assertHelper(0, false, queue.Back)

	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assertHelper(1, true, queue.Front)
	assertHelper(3, true, queue.Back)

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assertHelper(2, true, queue.Front)
	assertHelper(4, true, queue.Back)

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
