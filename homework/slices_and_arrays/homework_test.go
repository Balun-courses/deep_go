package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Number interface {
	int8 | int16 | int32 | int64
}

type CircularQueue[T Number] struct {
	size   int
	begin  int
	end    int
	values []T
}

func NewCircularQueue[T Number](size int) CircularQueue[T] {
	return CircularQueue[T]{
		size:   size,
		begin:  -1,
		end:    -1,
		values: make([]T, size),
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.begin != -1 && q.begin == q.end {
		return false
	}

	if q.begin == -1 {
		q.begin = 0
		q.end = 0
	}

	q.values[q.end] = value
	q.end = (q.end + 1) % q.size

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.begin == -1 {
		return false
	}

	q.begin = (q.begin + 1) % q.size

	if q.begin == q.end {
		q.begin = -1
		q.end = -1
	}

	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.begin != -1 {
		return q.values[q.begin]
	}

	return -1
}

func (q *CircularQueue[T]) Back() T {
	if q.begin != -1 {
		return q.values[(q.end+q.size-1)%q.size]
	}

	return -1
}

func (q *CircularQueue[T]) Empty() bool {
	return q.begin == -1
}

func (q *CircularQueue[T]) Full() bool {
	return q.begin != -1 && q.begin == q.end
}

func TestCircularQueueInt64(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int64](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, int64(-1), queue.Front())
	assert.Equal(t, int64(-1), queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int64{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, int64(1), queue.Front())
	assert.Equal(t, int64(3), queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int64{4, 2, 3}, queue.values))

	assert.Equal(t, int64(2), queue.Front())
	assert.Equal(t, int64(4), queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}

func TestCircularQueueInt8(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int8](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, int8(-1), queue.Front())
	assert.Equal(t, int8(-1), queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int8{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, int8(1), queue.Front())
	assert.Equal(t, int8(3), queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int8{4, 2, 3}, queue.values))

	assert.Equal(t, int8(2), queue.Front())
	assert.Equal(t, int8(4), queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
