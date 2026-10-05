package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
var NILValue = 0

type CircularQueue struct {
	values []int
	front  int
	back   int
}

func NewCircularQueue(size int) CircularQueue {
	slice := make([]int, 3)
	return CircularQueue{slice, -1, -1}
}

func (q *CircularQueue) Push(value int) bool {
	switch {
	case q.Empty():
		q.back = 0
		q.front = 0
		q.values[0] = value
		return true
	case q.Full():
		return false
	case q.back+1 == len(q.values):
		q.back = (q.back + 1) % len(q.values)
	default:
		q.back += 1
	}

	q.values[q.back] = value
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.values[q.front] = NILValue

	switch {
	case q.back == q.front:
		q.back = -1
		q.front = -1
	case q.front+1 == len(q.values):
		q.front = (q.front + 1) % len(q.values)
	default:
		q.front += 1
	}

	return true
}

func (q *CircularQueue) Front() int {
	if q.front == -1 {
		return q.front
	}
	return q.values[q.front]
}

func (q *CircularQueue) Back() int {
	if q.back == -1 {
		return q.back
	}
	return q.values[q.back]
}

func (q *CircularQueue) Empty() bool {
	if q.front == -1 && q.back == -1 {
		return true
	}
	return false
}

func (q *CircularQueue) Full() bool {
	if (q.front == 0 && q.back == len(q.values)-1) ||
		(q.back+1 == q.front) {
		return true
	}
	return false
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
