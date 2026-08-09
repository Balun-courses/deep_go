package main

import (
	"errors"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type COWBuffer struct {
	data   []byte
	refs   *int
	mu     *sync.Mutex
	closed bool
}

func COWFinalizerSetup(cb *COWBuffer) {
	runtime.SetFinalizer(cb, func(c *COWBuffer) {
		if !c.closed {
			c.Close()
		}
	})
}

func NewCOWBuffer(data []byte) *COWBuffer {
	cb := &COWBuffer{data: data, refs: new(1), mu: new(sync.Mutex)}

	COWFinalizerSetup(cb)

	return cb
}

func (b *COWBuffer) Clone() (*COWBuffer, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return nil, errors.New("buffer closed")
	}

	*b.refs++
	cb := &COWBuffer{data: b.data, refs: b.refs, mu: b.mu}
	COWFinalizerSetup(cb)

	return cb, nil
}

func (b *COWBuffer) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.closed = true
	b.data = nil
	if *b.refs > 0 {
		*b.refs--
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.data == nil || index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs <= 1 {
		b.data[index] = value
		return true
	}

	*b.refs--

	newData := make([]byte, len(b.data))
	copy(newData, b.data)

	b.data = newData
	b.refs = new(1)
	b.mu = new(sync.Mutex)

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}

func TestCOWBuffer(t *testing.T) {
	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1, err := buffer.Clone()
	assert.NoError(t, err)
	copy2, err := buffer.Clone()
	assert.NoError(t, err)

	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(copy1.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy1.String())) == unsafe.StringData(copy2.String()))

	assert.True(t, buffer.Update(0, 'g'))
	assert.False(t, buffer.Update(-1, 'g'))
	assert.False(t, buffer.Update(4, 'g'))

	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))

	assert.NotEqual(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	copy1.Close()
	_, err = copy1.Clone()
	assert.Error(t, err)

	previous := copy2.data
	copy2.Update(0, 'f')
	current := copy2.data

	// 1 reference - don't need to copy buffer during update
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))

	copy2.Close()
}

func TestCOWBufferFinalizer(t *testing.T) {
	buffer := NewCOWBuffer([]byte{'a', 'b', 'c', 'd'})
	defer buffer.Close()

	dropped, err := buffer.Clone()
	assert.NoError(t, err)
	assert.NotNil(t, dropped)
	assert.Equal(t, 2, *buffer.refs)

	dropped = nil

	for i := 0; i < 3; i++ {
		buffer.mu.Lock()
		refs := *buffer.refs
		buffer.mu.Unlock()
		if refs == 1 {
			break
		}
		runtime.GC()
		runtime.Gosched()
	}

	buffer.mu.Lock()
	assert.Equal(t, 1, *buffer.refs)
	buffer.mu.Unlock()

	previous := buffer.data
	assert.True(t, buffer.Update(0, 'g'))
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(buffer.data))
}

func TestCOWBufferConcurrent(t *testing.T) {
	buffer := NewCOWBuffer([]byte{'a', 'b', 'c', 'd'})
	defer buffer.Close()

	const workers = 16
	var wg sync.WaitGroup
	for i := range workers {
		wg.Add(1)
		go func(n byte) {
			defer wg.Done()
			cl, err := buffer.Clone()
			if err != nil {
				t.Error(err)
				return
			}
			defer cl.Close()
			for range 100 {
				if !cl.Update(0, n) {
					t.Error("update failed")
					return
				}
				_ = cl.String()
			}
		}(byte(i))
	}
	wg.Wait()
}
