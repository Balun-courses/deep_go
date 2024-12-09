package mysql

import "golang_course/lessons/interfaces/producer_interface/storage"

type MySQLStorage struct{}

func (s *MySQLStorage) GetAllClients() ([]storage.Client, error) {
	// already implemented
	return nil, nil
}

func (s *MySQLStorage) GetClientsByAge(age int) ([]storage.Client, error) {
	// already implemented
	return nil, nil
}

func (s *MySQLStorage) GetClient(id int) (storage.Client, error) {
	// already implemented
	return storage.Client{}, nil
}

func (s *MySQLStorage) RemoveClient(id int) error {
	// already implemented
	return nil
}

func (s *MySQLStorage) UpdateClient(client storage.Client) error {
	// already implemented
	return nil
}

func (s *MySQLStorage) CreateClient(client storage.Client) error {
	// already implemented
	return nil
}
