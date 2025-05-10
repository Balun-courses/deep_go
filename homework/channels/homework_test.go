package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
type Task struct {
	job          func()
	numOfRepeats int
}

type WorkerPool struct {
	tasks chan Task
	wg    sync.WaitGroup
}

func NewWorkerPool(workersNumber int) *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan Task, workersNumber),
		wg:    sync.WaitGroup{},
	}

	wp.wg.Add(workersNumber)

	for i := 0; i < workersNumber; i++ {
		go func() {
			defer wp.wg.Done()
			for task := range wp.tasks {
				for j := 0; j < task.numOfRepeats; j++ {
					task.job()
				}
			}
		}()
	}

	return wp
}

// Return an error if the pool is full
func (wp *WorkerPool) AddTask(task Task) error {
	select {
	case wp.tasks <- task:
		return nil
	default:
		go func() {
			wp.tasks <- task
		}()
		return fmt.Errorf("worker is full")
	}
}

// Shutdown all workers and wait for all
// tasks in the pool to complete
func (wp *WorkerPool) Shutdown() {
	close(wp.tasks)
	wp.wg.Wait()
}

func TestWorkerPool(t *testing.T) {
	var counter atomic.Int32
	task := func() {
		time.Sleep(time.Millisecond * 500)
		counter.Add(1)
	}

	pool := NewWorkerPool(2)
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})

	time.Sleep(time.Millisecond * 600)
	assert.Equal(t, int32(2), counter.Load())

	time.Sleep(time.Millisecond * 600)
	assert.Equal(t, int32(3), counter.Load())

	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	pool.Shutdown() // wait tasks

	assert.Equal(t, int32(6), counter.Load())
}

func TestWorkerPoolWithRepeats(t *testing.T) {
	var counter atomic.Int32
	task := func() {
		time.Sleep(time.Millisecond * 500)
		counter.Add(1)
	}

	pool := NewWorkerPool(2)
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 3,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 3,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})

	time.Sleep(time.Millisecond * 600 * 3)
	assert.Equal(t, int32(6), counter.Load())

	time.Sleep(time.Millisecond * 600)
	assert.Equal(t, int32(7), counter.Load())

	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 2,
	})
	_ = pool.AddTask(Task{
		job:          task,
		numOfRepeats: 1,
	})
	pool.Shutdown() // wait tasks

	assert.Equal(t, int32(11), counter.Load())
}
