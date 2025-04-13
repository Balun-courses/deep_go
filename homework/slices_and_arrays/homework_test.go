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
	begin  int
	tail   int
	size   int
	// need to implement
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		size:   size,
		begin:  -1,
	} // need to implement
}

func (q *CircularQueue) Push(value int) bool {
	if q.tail == q.begin {
		return false
	}

	q.values[q.tail] = value
	if q.begin == -1 {
		q.begin = q.tail
	}

	if q.tail+1 >= q.size {
		q.tail = 0
	} else {
		q.tail++
	}

	return true // need to implement
}

func (q *CircularQueue) Pop() bool {
	if q.begin == -1 {
		return false
	}

	q.begin++
	if q.begin >= q.size {
		q.begin = 0
	}

	if q.tail == q.begin {
		q.begin = -1
	}

	return true // need to implement
}

func (q *CircularQueue) Front() int {
	if q.begin == -1 {
		return -1
	}

	return q.values[q.begin] // need to implement
}

func (q *CircularQueue) Back() int {
	if q.begin == -1 {
		return -1
	}

	pre := q.tail - 1
	if pre < 0 {
		pre = q.size - 1
	}

	return q.values[pre] // need to implement
}

func (q *CircularQueue) Empty() bool {
	return q.begin == -1 // need to implement
}

func (q *CircularQueue) Full() bool {
	return q.begin == q.tail // need to implement
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
	fmt.Println(queue.values)
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
