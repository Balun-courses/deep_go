package service

import "golang_course/lessons/interfaces/producer_interface/storage"

type ClientUpdater struct {
	repository storage.ClientStorage
}

func NewClientUpdater(repository storage.ClientStorage) ClientUpdater {
	return ClientUpdater{
		repository: repository,
	}
}

func (s *ClientUpdater) UpdateClient(client storage.Client) error {
	return s.repository.UpdateClient(client)
}
