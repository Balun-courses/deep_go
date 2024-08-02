package main

type Cient struct{}
type Admin struct{}

type Storage struct{}

func (s *Storage) Get(id int) (interface{}, error) {
	// already implemented
	return nil, nil
}

func (s *Storage) Set(id int, user interface{}) error {
	// already implemented
	return nil
}
