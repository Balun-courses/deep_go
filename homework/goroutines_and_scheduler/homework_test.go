package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	tasks   []Task
	taskMap map[int]int // Отображение ID задачи на индекс в куче
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		taskMap: make(map[int]int),
	}
}

func (s *Scheduler) AddTask(task Task) {
	if _, exists := s.taskMap[task.Identifier]; exists {
		return
	}

	s.tasks = append(s.tasks, task)
	index := len(s.tasks) - 1
	s.taskMap[task.Identifier] = index

	s.siftUp(index)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	index, exists := s.taskMap[taskID]
	if !exists {
		return
	}

	oldPriority := s.tasks[index].Priority
	s.tasks[index].Priority = newPriority

	if newPriority > oldPriority {
		s.siftUp(index)
	} else if newPriority < oldPriority {
		s.siftDown(index)
	}
}

func (s *Scheduler) GetTask() Task {
	if len(s.tasks) == 0 {
		return Task{Identifier: -1}
	}

	maxTask := s.tasks[0]
	delete(s.taskMap, maxTask.Identifier)

	lastIndex := len(s.tasks) - 1
	if lastIndex > 0 {
		s.tasks[0] = s.tasks[lastIndex]
		s.taskMap[s.tasks[0].Identifier] = 0
	}
	s.tasks = s.tasks[:lastIndex]

	if len(s.tasks) > 0 {
		s.siftDown(0)
	}

	return maxTask
}

func (s *Scheduler) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if s.tasks[parent].Priority >= s.tasks[index].Priority {
			break
		}
		s.swap(index, parent)
		index = parent
	}
}

func (s *Scheduler) siftDown(index int) {
	size := len(s.tasks)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < size && s.tasks[left].Priority > s.tasks[largest].Priority {
			largest = left
		}
		if right < size && s.tasks[right].Priority > s.tasks[largest].Priority {
			largest = right
		}
		if largest == index {
			break
		}
		s.swap(index, largest)
		index = largest
	}
}

func (s *Scheduler) swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
	s.taskMap[s.tasks[i].Identifier] = i
	s.taskMap[s.tasks[j].Identifier] = j
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task := scheduler.GetTask()
	assert.Equal(t, task5, task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)

	task = scheduler.GetTask()
	assert.Equal(t, task1.Identifier, task.Identifier)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}
