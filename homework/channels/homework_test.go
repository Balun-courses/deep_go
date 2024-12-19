package main

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type WorkerPool struct {
	// need to implement
}

func NewWorkerPool(workersNumber int) *WorkerPool {
	// need to implement
	return &WorkerPool{}
}

// Return an error if the pool is full
func (wp *WorkerPool) AddTask(task func()) error {
	// need to implement
	return nil
}

// Shutdown all workers and wait for all
// tasks in the pool to complete
func (wp *WorkerPool) Shutdown() {
	// need to implement
}

func TestWorkerPool(t *testing.T) {
	var counter atomic.Int32
	task := func() {
		time.Sleep(time.Millisecond * 500)
		counter.Add(1)
	}

	pool := NewWorkerPool(2)
	_ = pool.AddTask(task)
	_ = pool.AddTask(task)
	_ = pool.AddTask(task)

	time.Sleep(time.Millisecond * 600)
	assert.Equal(t, int32(2), counter.Load())

	time.Sleep(time.Millisecond * 600)
	assert.Equal(t, int32(3), counter.Load())

	_ = pool.AddTask(task)
	_ = pool.AddTask(task)
	_ = pool.AddTask(task)
	pool.Shutdown() // wait tasks

	assert.Equal(t, int32(6), counter.Load())
}
