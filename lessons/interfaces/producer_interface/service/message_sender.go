package service

import "golang_course/lessons/interfaces/producer_interface/storage"

type MessageSender struct {
	repository storage.ClientStorage
}

func NewMessageSender(repository storage.ClientStorage) MessageSender {
	return MessageSender{
		repository: repository,
	}
}

func (s *MessageSender) SendMessage(userId int, message string) error {
	_, err := s.repository.GetClient(userId)
	if err != nil {
		return err
	}

	// send message to client

	return nil
}
