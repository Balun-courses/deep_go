package main

import "sync"

type Storage struct {
	sync.Mutex
	data map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Get(key string) (string, bool) {
	s.Lock()
	defer s.Unlock()

	value, found := s.data[key]
	return value, found
}

func main() {
	storage := NewStorage()
	storage.Get("key_1")

	storage.Lock()   // dangerous
	storage.Unlock() // dangerous
}
