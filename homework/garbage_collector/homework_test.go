package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

// go test -v homework_test.go

type Node struct {
	next *Node
	val  int
}

func NewSeenMap(mapSize int) *SeenMap {
	return &SeenMap{
		data: make(map[uintptr]struct{}, mapSize),
		mx:   sync.RWMutex{},
	}
}

type SeenMap struct {
	data map[uintptr]struct{}
	mx   sync.RWMutex
}

func (s *SeenMap) Exists(key uintptr) bool {
	s.mx.RLock()
	_, ok := s.data[key]
	s.mx.RUnlock()

	return ok
}

func (s *SeenMap) Add(key uintptr) {
	s.mx.Lock()
	s.data[key] = struct{}{}
	s.mx.Unlock()
}

func Trace(stacks [][]uintptr) []uintptr {
	if len(stacks) == 0 {
		return []uintptr{}
	}

	res := make([]uintptr, 0)
	seenMap := make(map[uintptr]struct{})

	var tracePtr func(ptr uintptr)
	tracePtr = func(ptr uintptr) {
		if ptr == 0 {
			return
		}

		if _, ok := seenMap[ptr]; ok {
			return
		}

		seenMap[ptr] = struct{}{}
		res = append(res, ptr)

		val := (*unsafe.Pointer)(unsafe.Pointer(ptr))
		if val != nil && uintptr(*val) != 0 {
			tracePtr(uintptr(*val))
		}
	}

	for _, stack := range stacks {
		for _, ptr := range stack {
			tracePtr(ptr)
		}
	}

	return res
}

func TraceParallel(stacks [][]uintptr) []uintptr {
	numGoroutines := len(stacks)
	wg := &sync.WaitGroup{}
	wg.Add(numGoroutines)
	seenMap := NewSeenMap(numGoroutines)
	res := make([]uintptr, 0)
	mx := sync.Mutex{}

	var tracePtr func(ptr uintptr)
	tracePtr = func(ptr uintptr) {
		if ptr == 0 {
			return
		}

		if seenMap.Exists(ptr) {
			return
		}

		seenMap.Add(ptr)

		mx.Lock()
		res = append(res, ptr)
		mx.Unlock()

		val := (*unsafe.Pointer)(unsafe.Pointer(ptr))
		if val != nil {
			tracePtr(uintptr(*val))
		}
	}

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < len(stacks[i]); j++ {
				tracePtr(stacks[i][j])
			}

		}()
	}

	wg.Wait()

	return res
}

func TestTrace(t *testing.T) {
	t.Run("No cycles", func(t *testing.T) {
		var heapObjects = []int{
			0x00, 0x00, 0x00, 0x00, 0x00,
		}

		var heapPointer1 *int = &heapObjects[1]
		var heapPointer2 *int = &heapObjects[2]
		var heapPointer3 *int = nil
		var heapPointer4 **int = &heapPointer3

		var stacks = [][]uintptr{
			{
				uintptr(unsafe.Pointer(&heapPointer1)), 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[0])),
				0x00, 0x00, 0x00, 0x00,
			},
			{
				uintptr(unsafe.Pointer(&heapPointer2)), 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[1])),
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[2])),
				uintptr(unsafe.Pointer(&heapPointer4)), 0x00, 0x00, 0x00,
			},
			{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[3])),
			},
		}

		pointers := Trace(stacks)
		expectedPointers := []uintptr{
			uintptr(unsafe.Pointer(&heapPointer1)),
			uintptr(unsafe.Pointer(&heapObjects[1])),
			uintptr(unsafe.Pointer(&heapObjects[0])),
			uintptr(unsafe.Pointer(&heapPointer2)),
			uintptr(unsafe.Pointer(&heapObjects[2])),
			uintptr(unsafe.Pointer(&heapPointer4)),
			uintptr(unsafe.Pointer(&heapPointer3)),
			uintptr(unsafe.Pointer(&heapObjects[3])),
		}

		assert.True(t, reflect.DeepEqual(pointers, expectedPointers))
	})

	t.Run("With cycle", func(t *testing.T) {
		node1 := &Node{
			val: 1,
		}
		node2 := &Node{
			val: 2,
		}
		node3 := &Node{
			val:  3,
			next: node1,
		}
		node1.next = node2
		node2.next = node3

		var heapObjects = []int{
			0x00, 0x00, 0x00, 0x00, 0x00,
		}

		var heapPointer1 *int = &heapObjects[1]
		var heapPointer2 *int = &heapObjects[2]
		var heapPointer3 *int = nil
		var heapPointer4 **int = &heapPointer3

		var stacks = [][]uintptr{
			{
				uintptr(unsafe.Pointer(&heapPointer1)), uintptr(unsafe.Pointer(&node1)), 0x00, 0x00,
				uintptr(unsafe.Pointer(&node2)), 0x00, uintptr(unsafe.Pointer(&node3)), 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[0])),
				0x00, 0x00, 0x00, 0x00,
			},
			{
				uintptr(unsafe.Pointer(&heapPointer2)), 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[1])),
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[2])),
				uintptr(unsafe.Pointer(&heapPointer4)), 0x00, 0x00, 0x00,
			},
			{
				uintptr(unsafe.Pointer(&node1.next)), 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[3])),
			},
		}

		pointers := Trace(stacks)
		expectedPointers := []uintptr{
			uintptr(unsafe.Pointer(&heapPointer1)),
			uintptr(unsafe.Pointer(&heapObjects[1])),
			uintptr(unsafe.Pointer(&node1)),
			uintptr(unsafe.Pointer(&node1.next)),
			uintptr(unsafe.Pointer(&node2.next)),
			uintptr(unsafe.Pointer(&node3.next)),
			uintptr(unsafe.Pointer(&node2)),
			uintptr(unsafe.Pointer(&node3)),
			uintptr(unsafe.Pointer(&heapObjects[0])),
			uintptr(unsafe.Pointer(&heapPointer2)),
			uintptr(unsafe.Pointer(&heapObjects[2])),
			uintptr(unsafe.Pointer(&heapPointer4)),
			uintptr(unsafe.Pointer(&heapPointer3)),
			uintptr(unsafe.Pointer(&heapObjects[3])),
		}

		assert.True(t, reflect.DeepEqual(pointers, expectedPointers))
	})

	t.Run("one node tracing", func(t *testing.T) {
		node1 := &Node{
			val: 1,
		}
		node2 := &Node{
			val: 2,
		}
		node3 := &Node{
			val:  3,
			next: node1,
		}
		node1.next = node2
		node2.next = node3

		var stacks = [][]uintptr{
			{
				uintptr(unsafe.Pointer(&node1)),
			},
		}

		pointers := Trace(stacks)
		expectedPointers := []uintptr{
			uintptr(unsafe.Pointer(&node1)),
			uintptr(unsafe.Pointer(&node1.next)),
			uintptr(unsafe.Pointer(&node2.next)),
			uintptr(unsafe.Pointer(&node3.next)),
		}

		assert.True(t, reflect.DeepEqual(pointers, expectedPointers))
	})
}
