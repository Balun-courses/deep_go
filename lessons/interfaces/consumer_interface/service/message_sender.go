package service

import "golang_course/lessons/interfaces/consumer_interface/entity"

type clientGetter interface {
	GetClient(int) (entity.Client, error)
}

type MessageSender struct {
	repository clientGetter
}

func NewMessageSender(repository clientGetter) MessageSender {
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
