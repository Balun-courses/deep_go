package main

import "fmt"

type Set[K comparable] struct {
	data map[K]struct{}
}

func NewSet[K comparable]() Set[K] {
	return Set[K]{
		data: make(map[K]struct{}),
	}
}

func (s *Set[K]) Insert(key K) {
	s.data[key] = struct{}{}
}

func (s *Set[K]) Erase(key K) {
	delete(s.data, key)
}

func (s *Set[K]) Contains(key K) bool {
	_, found := s.data[key]
	return found
}

// skipping like with function
func (s *Set[_]) Print() {
	fmt.Println(s.data)
}

func main() {
	set := NewSet[string]()
	set.Insert("key")
	set.Erase("key")
}
