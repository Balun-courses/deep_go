package redis

import "golang_course/lessons/interfaces/producer_interface/storage"

type RedisStorage struct{}

func (s *RedisStorage) GetAllClients() ([]storage.Client, error) {
	// already implemented
	return nil, nil
}

func (s *RedisStorage) GetClientsByAge(age int) ([]storage.Client, error) {
	// already implemented
	return nil, nil
}

func (s *RedisStorage) GetClient(id int) (storage.Client, error) {
	// already implemented
	return storage.Client{}, nil
}

func (s *RedisStorage) RemoveClient(id int) error {
	// already implemented
	return nil
}

func (s *RedisStorage) UpdateClient(client storage.Client) error {
	// already implemented
	return nil
}

func (s *RedisStorage) CreateClient(client storage.Client) error {
	// already implemented
	return nil
}
