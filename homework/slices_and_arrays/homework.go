package main

type CircularQueue struct {
	values      []int
	head, tail  int
	size, count int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		head:   0,
		tail:   0,
		size:   size,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.tail] = value
	q.tail = (q.tail + 1) % q.size
	q.count++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % q.size
	q.count--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.head]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[(q.tail-1+q.size)%q.size]
}

func (q *CircularQueue) Empty() bool {
	if q.Full() {
		return false
	}
	return q.count == 0
}

func (q *CircularQueue) Full() bool {
	return q.count == q.size
}

//на дженериках

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type CircularQueueGen[T Integer] struct {
	values      []T
	head, tail  int
	size, count int
}

func NewCircularQueueGen[T Integer](size int) *CircularQueueGen[T] {
	return &CircularQueueGen[T]{
		values: make([]T, size),
		size:   size,
		head:   0,
		tail:   0,
	}
}

func (q *CircularQueueGen[T]) Push(value T) bool {
	if q.Full() {
		return false
	}
	q.values[q.tail] = value
	q.tail = (q.tail + 1) % q.size
	q.count++
	return true
}

func (q *CircularQueueGen[T]) Pop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % q.size
	q.count--
	return true
}

func (q *CircularQueueGen[T]) Front() T {
	var zero T
	if q.Empty() {
		return zero
	}
	return q.values[q.head]
}

func (q *CircularQueueGen[T]) Back() T {
	var zero T
	if q.Empty() {
		return zero
	}
	return q.values[(q.tail-1+q.size)%q.size]
}

func (q *CircularQueueGen[T]) Empty() bool {
	return q.count == 0
}

func (q *CircularQueueGen[T]) Full() bool {
	return q.count == q.size
}
