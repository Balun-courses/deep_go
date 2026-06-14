package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue[T any] struct {
	values     []T
	readIndex  int
	writeIndex int
	size       int
}

func NewCircularQueue[T any](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values:     make([]T, size),
		readIndex:  0,
		writeIndex: 0,
		size:       0,
	}
}

func (queue *CircularQueue[T]) Push(value T) bool {
	if queue.Full() {
		return false
	}

	queue.values[queue.readIndex] = value
	queue.readIndex = (queue.readIndex + 1) % len(queue.values)
	queue.size++

	return true
}

func (queue *CircularQueue[T]) Pop() bool {
	if queue.Empty() {
		return false
	}

	// Обнуляем на ZeroValue, чтобы GC мог обработать ссылочные типы
	var zero T
	queue.values[queue.writeIndex] = zero

	queue.writeIndex = (queue.writeIndex + 1) % len(queue.values)
	queue.size--

	return true
}

func (queue *CircularQueue[T]) Front() T {
	if queue.Empty() {
		var zero T
		return zero
	}

	return queue.values[queue.writeIndex]
}

func (queue *CircularQueue[T]) Back() T {
	if queue.Empty() {
		var zero T
		return zero
	}
	last := (queue.readIndex - 1 + len(queue.values)) % len(queue.values)
	return queue.values[last]
}

func (queue *CircularQueue[T]) Empty() bool {
	return queue.size == 0
}

func (queue *CircularQueue[T]) Full() bool {
	return len(queue.values) == queue.size
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	// Так как делаю реализацию через дженерики, осознанно поменял на NullValue для int
	assert.Equal(t, 0, queue.Front())
	assert.Equal(t, 0, queue.Back())
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
