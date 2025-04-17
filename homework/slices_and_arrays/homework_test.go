package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values   []int
	popIdx   uint
	insIdx   uint
	currSize int
}

func NewCircularQueue(size int) CircularQueue {
	c := CircularQueue{
		values: make([]int, size),
	}
	for i := range c.values {
		c.values[i] = -1
	}
	return c
}

func (q *CircularQueue) Push(value int) bool {
	if q.currSize == len(q.values) {
		return false
	}
	q.currSize++
	q.values[q.insIdx%uint(len(q.values))] = value
	q.insIdx++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.currSize == 0 {
		return false
	}
	q.currSize--
	q.values[q.popIdx%uint(len(q.values))] = -1
	q.popIdx++
	return true
}

func (q *CircularQueue) Front() int {
	return q.values[q.popIdx%uint(len(q.values))]
}

func (q *CircularQueue) Back() int {
	idx := q.insIdx - 1
	return q.values[idx%uint(len(q.values))]
}

func (q *CircularQueue) Empty() bool {
	return q.currSize == 0
}

func (q *CircularQueue) Full() bool {
	return q.currSize == len(q.values)
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
