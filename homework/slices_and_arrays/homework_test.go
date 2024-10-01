package main

import (
	"testing"
)

// go test -v homework_test.go

type Deque struct {
	// need to implement
}

func NewDeque() Deque {
	return Deque{}
}

func (d *Deque) PushFront(value int) {
	// need to implement
}

func (d *Deque) PushBack(value int) {
	// need to implement
}

func (d *Deque) PopFront() bool {
	return false // need to implement
}

func (d *Deque) PopBack() bool {
	return false // need to implement
}

func (d *Deque) Front() (int, bool) {
	return 0, false // need to implement
}

func (d *Deque) Back() (int, bool) {
	return 0, false // need to implement
}

func (d *Deque) At(index int) (int, bool) {
	return 0, false // need to implement
}

func (d *Deque) Size() int {
	return 0 // need to implement
}

func (d *Deque) ForEach(action func(int)) {
	// need to implement
}

func TestDeque(t *testing.T) {

}
