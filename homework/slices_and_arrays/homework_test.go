package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values  []int
	size    int
	headIdx *int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
	}
}

func (q *CircularQueue) Push(value int) bool {
	// queue is full
	if q.Full() {
		return false
	}

	// get next free position index
	nextFreeIdx := func() int {
		// queue is empty
		if q.headIdx == nil {
			return 0
		}
		return (*q.headIdx + q.size) % len(q.values)
	}()

	// push value
	q.values[nextFreeIdx] = value
	q.size++

	// set head if it is the first value
	if q.headIdx == nil {
		q.headIdx = &nextFreeIdx
	}

	return true
}

func (q *CircularQueue) Pop() bool {
	// empty queue
	if q.Empty() {
		return false
	}

	q.size--

	// empty queue
	if q.Empty() {
		q.headIdx = nil
		return true
	}

	// shift queue head by 1
	*q.headIdx = (*q.headIdx + 1) % len(q.values)

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[*q.headIdx]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	lastValueIdx := (*q.headIdx + q.size - 1) % len(q.values)
	return q.values[lastValueIdx]
}

func (q *CircularQueue) Empty() bool {
	return q.size == 0
}

func (q *CircularQueue) Full() bool {
	return q.size == len(q.values)
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
