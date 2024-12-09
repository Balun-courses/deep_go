package mysql

import "golang_course/lessons/interfaces/consumer_interface/entity"

type MySQLStorage struct{}

func (s *MySQLStorage) GetAllClients() ([]entity.Client, error) {
	// already implemented
	return nil, nil
}

func (s *MySQLStorage) GetClientsByAge(age int) ([]entity.Client, error) {
	// already implemented
	return nil, nil
}

func (s *MySQLStorage) GetClient(id int) (entity.Client, error) {
	// already implemented
	return entity.Client{}, nil
}

func (s *MySQLStorage) RemoveClient(id int) error {
	// already implemented
	return nil
}

func (s *MySQLStorage) UpdateClient(client entity.Client) error {
	// already implemented
	return nil
}

func (s *MySQLStorage) CreateClient(client entity.Client) error {
	// already implemented
	return nil
}
