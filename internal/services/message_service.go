package services

import "2task/internal/repositories"

type MessageRepository = repositories.MessageRepository
type Message = repositories.Message

type MessageService struct {
	repo MessageRepository
}

func NewMessageService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) UpdateMessageByID(id int, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

func (s *MessageService) DeleteMessageByID(id int) error {
	return s.repo.DeleteMessageByID(id)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}
