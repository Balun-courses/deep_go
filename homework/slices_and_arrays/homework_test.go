package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
type Number interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T Number] struct {
	values []T
	front  int
	rear   int
}

func NewCircularQueue[T Number](size int) CircularQueue[T] {
	if size <= 0 {
		return CircularQueue[T]{}
	}
	v := make([]T, size)
	return CircularQueue[T]{values: v, front: -1, rear: -1}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}
	if q.front == -1 {
		q.front = 0
	}
	q.rear = (q.rear + 1) % len(q.values)
	q.values[q.rear] = value
	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		return true
	}
	q.front = (q.front + 1) % len(q.values)
	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.front == -1
}

func (q *CircularQueue[T]) Full() bool {
	return (q.front == 0 && q.rear+1 == len(q.values)) || (q.rear == (q.front-1)%len(q.values))
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))
	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
