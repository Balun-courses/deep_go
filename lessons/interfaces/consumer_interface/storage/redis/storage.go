package redis

import "golang_course/lessons/interfaces/consumer_interface/entity"

type RedisStorage struct{}

func (s *RedisStorage) GetAllClients() ([]entity.Client, error) {
	// already implemented
	return nil, nil
}

func (s *RedisStorage) GetClientsByAge(age int) ([]entity.Client, error) {
	// already implemented
	return nil, nil
}

func (s *RedisStorage) GetClient(id int) (entity.Client, error) {
	// already implemented
	return entity.Client{}, nil
}

func (s *RedisStorage) RemoveClient(id int) error {
	// already implemented
	return nil
}

func (s *RedisStorage) UpdateClient(client entity.Client) error {
	// already implemented
	return nil
}

func (s *RedisStorage) CreateClient(client entity.Client) error {
	// already implemented
	return nil
}
