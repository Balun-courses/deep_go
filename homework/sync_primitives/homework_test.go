package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type RWMutex struct {
	c            *sync.Cond
	readersCount int64
	hasWriter    bool
	state        int32
}

const unlocked = 0
const locked = 1

func NewRWMutex() *RWMutex {
	return &RWMutex{sync.NewCond(new(sync.Mutex)), 0, false, 0}
}

func (m *RWMutex) Lock() {
	m.c.L.Lock()

	for m.hasWriter {
		m.c.Wait()
	}

	m.hasWriter = true
	m.state = locked

	for m.readersCount > 0 {
		m.c.Wait()
	}

	m.c.L.Unlock()
}

func (m *RWMutex) TryLock() bool {
	old := m.state
	if old == locked {
		return false
	}

	if !atomic.CompareAndSwapInt32(&m.state, old, locked) {
		return false
	}

	m.hasWriter = true

	return true
}

func (m *RWMutex) TryRLock() bool {
	if m.hasWriter {
		return false
	}

	// m.hasWriter == true ?
	// если нет читателей и писателей, но не удалось поменять статус, то false
	if atomic.LoadInt64(&m.readersCount) == 0 && !atomic.CompareAndSwapInt32(&m.state, unlocked, locked) {
		return false
	}

	m.readersCount++

	return true
}

func (m *RWMutex) Unlock() {
	m.c.L.Lock()
	m.hasWriter = false
	m.state = unlocked
	m.c.Broadcast()
	m.c.L.Unlock()
}

func (m *RWMutex) RLock() {
	m.c.L.Lock()

	for m.hasWriter {
		m.c.Wait()
	}

	m.state = locked
	m.readersCount++
	m.c.L.Unlock()
}

func (m *RWMutex) RUnlock() {
	m.c.L.Lock()
	m.readersCount--
	m.state = unlocked

	if m.readersCount == 0 {
		m.c.Broadcast()
	}

	m.c.L.Unlock()
}

func TestRWMutexWithWriter(t *testing.T) {
	mutex := NewRWMutex()
	mutex.Lock() // writer

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)
	var mutualExlusionWithReader atomic.Bool
	mutualExlusionWithReader.Store(true)

	go func() {
		mutex.Lock() // another writer
		assert.False(t, mutex.TryLock())
		assert.False(t, mutex.TryRLock())
		mutualExlusionWithWriter.Store(false)
	}()

	go func() {
		mutex.RLock() // another reader
		assert.True(t, mutex.TryRLock())
		mutex.RUnlock()
		mutualExlusionWithReader.Store(false)
	}()

	time.Sleep(time.Second)
	assert.True(t, mutualExlusionWithWriter.Load())
	assert.True(t, mutualExlusionWithReader.Load())
}

func TestRWMutexWithReaders(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)

	go func() {
		mutex.Lock() // another writer
		mutualExlusionWithWriter.Store(false)
	}()

	time.Sleep(time.Second)
	assert.True(t, mutualExlusionWithWriter.Load())
}

func TestRWMutexMultipleReaders(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var readersCount atomic.Int32
	readersCount.Add(1)

	go func() {
		mutex.RLock() // another reader
		assert.True(t, mutex.TryRLock())
		readersCount.Add(2)
	}()

	go func() {
		mutex.RLock() // another reader
		assert.True(t, mutex.TryRLock())
		readersCount.Add(2)
	}()

	time.Sleep(time.Second)
	assert.Equal(t, int32(5), readersCount.Load())
}

func TestRWMutexWithWriterPriority(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)
	var readersCount atomic.Int32
	readersCount.Add(1)

	go func() {
		mutex.Lock() // another writer is waiting for reader
		mutualExlusionWithWriter.Store(false)
	}()

	time.Sleep(time.Second)

	go func() {
		mutex.RLock() // another reader is waiting for a higher priority writer
		readersCount.Add(1)
	}()

	go func() {
		mutex.RLock() // another reader is waiting for a higher priority writer
		readersCount.Add(1)
	}()

	time.Sleep(time.Second)

	assert.True(t, mutualExlusionWithWriter.Load())
	assert.Equal(t, int32(1), readersCount.Load())
}

type ReaderCountRWLock struct {
	m           sync.Mutex
	readerCount int64
}

func (l *ReaderCountRWLock) RLock() {
	//atomic.AddInt64(&l.readerCount, 1)
	l.m.Lock()
	l.readerCount++
	l.m.Unlock()
}

func (l *ReaderCountRWLock) RUnlock() {
	//atomic.AddInt64(&l.readerCount, -1)
	l.m.Lock()
	l.readerCount--
	l.m.Unlock()
}

func (l *ReaderCountRWLock) WUnlock() {
	l.m.Unlock()
}

func (l *ReaderCountRWLock) WLock() {
	for {
		l.m.Lock()
		if l.readerCount > 0 {
			l.m.Unlock()
		} else {
			break
		}
	}
}

func TestRWMutexWithWriterTest(t *testing.T) {
	mutex := ReaderCountRWLock{}
	mutex.WLock() // writer

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)
	var mutualExlusionWithReader atomic.Bool
	mutualExlusionWithReader.Store(true)

	go func() {
		mutex.WLock() // another writer
		mutualExlusionWithWriter.Store(false)
	}()

	go func() {
		mutex.RLock() // another reader
		mutualExlusionWithReader.Store(false)
	}()

	time.Sleep(time.Second)
	assert.True(t, mutualExlusionWithWriter.Load())
	assert.True(t, mutualExlusionWithReader.Load())
}
