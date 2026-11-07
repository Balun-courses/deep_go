package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// Trace обходит все объекты, достижимые из стеков.
// Каждый объект — одна машинная ячейка: если значение по адресу непустое,
// это указатель на следующий объект. Обход рекурсивный (DFS),
// уже посещённые объекты повторно не обходятся.
func Trace(stacks [][]uintptr) []uintptr {
	visited := make(map[uintptr]struct{})
	var roots []uintptr

	// Сначала собираем все корни со стеков (без дубликатов).
	// Корни считаются уже известными, чтобы не обходить их повторно,
	// если на них же укажет другой объект.
	for _, stack := range stacks {
		for _, slot := range stack {
			if slot == 0 {
				continue
			}
			if _, ok := visited[slot]; !ok {
				visited[slot] = struct{}{}
				roots = append(roots, slot)
			}
		}
	}

	result := make([]uintptr, 0, len(roots))

	var trace func(ptr uintptr)
	trace = func(ptr uintptr) {
		result = append(result, ptr)

		child := *(*uintptr)(unsafe.Pointer(ptr))
		if child == 0 {
			return
		}
		if _, ok := visited[child]; !ok {
			visited[child] = struct{}{}
			trace(child)
		}
	}

	for _, root := range roots {
		trace(root)
	}

	return result
}

func TestTrace(t *testing.T) {
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
		uintptr(unsafe.Pointer(&heapObjects[0])),
		uintptr(unsafe.Pointer(&heapPointer2)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
		uintptr(unsafe.Pointer(&heapObjects[2])),
		uintptr(unsafe.Pointer(&heapPointer4)),
		uintptr(unsafe.Pointer(&heapPointer3)),
		uintptr(unsafe.Pointer(&heapObjects[3])),
	}

	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestTraceDeepChain(t *testing.T) {
	var heapObjects = []int{0x00, 0x00, 0x00}

	var level1 *int = &heapObjects[0]
	var level2 **int = &level1
	var level3 ***int = &level2

	stacks := [][]uintptr{
		{uintptr(unsafe.Pointer(&level3)), 0x00, 0x00},
	}

	pointers := Trace(stacks)
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&level3)),
		uintptr(unsafe.Pointer(&level2)),
		uintptr(unsafe.Pointer(&level1)),
		uintptr(unsafe.Pointer(&heapObjects[0])),
	}

	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestTraceSharedObject(t *testing.T) {
	var heapObjects = []int{0x00, 0x00, 0x00}

	var shared *int = &heapObjects[1]
	var root1 **int = &shared
	var root2 **int = &shared

	stacks := [][]uintptr{
		{uintptr(unsafe.Pointer(&root1)), 0x00},
		{uintptr(unsafe.Pointer(&root2)), 0x00},
	}

	pointers := Trace(stacks)
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&root1)),
		uintptr(unsafe.Pointer(&shared)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
		uintptr(unsafe.Pointer(&root2)),
	}

	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestTraceChildIsAlsoRoot(t *testing.T) {
	var heapObjects = []int{0x00, 0x00}

	var p *int = &heapObjects[1]

	// объект &heapObjects[1] — ребёнок &p и одновременно корень второго стека
	stacks := [][]uintptr{
		{uintptr(unsafe.Pointer(&p)), 0x00},
		{0x00, uintptr(unsafe.Pointer(&heapObjects[1]))},
	}

	pointers := Trace(stacks)
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&p)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
	}

	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestTraceNoRoots(t *testing.T) {
	assert.Empty(t, Trace(nil))
	assert.Empty(t, Trace([][]uintptr{{0x00, 0x00}, {}}))
}
