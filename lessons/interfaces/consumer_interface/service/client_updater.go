package service

import "golang_course/lessons/interfaces/consumer_interface/entity"

type clientUpdater interface {
	UpdateClient(entity.Client) error
}

type ClientUpdater struct {
	repository clientUpdater
}

func NewClientUpdater(repository clientUpdater) ClientUpdater {
	return ClientUpdater{
		repository: repository,
	}
}

func (s *ClientUpdater) UpdateClient(client entity.Client) error {
	return s.repository.UpdateClient(client)
}
