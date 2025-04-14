package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values     []int
	size       int64
	start, end int64
}

func NewCircularQueue(size int) CircularQueue {
	values := make([]int, size)
	for i := range values {
		values[i] = -1
	}
	return CircularQueue{
		values: values,
		size:   int64(size),
		start:  0,
		end:    0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.values[q.end] != -1 {
		return false
	}

	q.values[q.end] = value
	q.end = (q.end + 1) % q.size

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.values[q.start] == -1 {
		return false
	}

	q.values[q.start] = -1
	q.start = (q.start + 1) % q.size

	return true
}

func (q *CircularQueue) Front() int {
	return q.values[q.start]
}

func (q *CircularQueue) Back() int {
	if q.end-1 < 0 {
		return q.values[q.size-1]
	}
	return q.values[q.end-1]
}

func (q *CircularQueue) Empty() bool {
	return (q.start == q.end) && (q.values[q.start] == -1)
}

func (q *CircularQueue) Full() bool {
	return (q.start == q.end) && (q.values[q.start] != -1)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

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
