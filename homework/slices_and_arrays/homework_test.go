package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	// need to implement
	start int
	end   int
}

func NewCircularQueue(size int) CircularQueue {
	if size <= 0 {
		fmt.Println("Default value for size will be applied")
		size = 10
	}
	return CircularQueue{
		values: make([]int, size),
		end:    -1,
	} // need to implement
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.start] = value
	q.start = (q.start + 1) % cap(q.values)
	if q.Empty() {
		q.end = q.start
	}
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.end == -1 {
		return false
	}
	q.end = (q.end + 1) % cap(q.values)
	if q.Full() {
		q.end = -1
	}
	return true
}

func (q *CircularQueue) Front() int {
	if q.end == -1 {
		return -1
	}
	v := (q.end - 1 + cap(q.values)) % cap(q.values)
	return q.values[v]
}

func (q *CircularQueue) Back() int {
	if q.end == -1 {
		return -1
	}
	v := (q.start - 1 + cap(q.values)) % cap(q.values)
	return q.values[v]
}

func (q *CircularQueue) Empty() bool {
	return q.end == -1
}

func (q *CircularQueue) Full() bool {
	if q.end != -1 && (q.start+1)%cap(q.values) == q.end {
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

func TestCircularQueueOne(t *testing.T) {
	const queueSize = 1
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.False(t, queue.Push(2))
	assert.False(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, queue.Pop())
	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(2))
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 2, queue.Back())
	assert.False(t, queue.Push(3))
}
