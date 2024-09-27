package repositories

import (
	"2task/internal/database/models"
	"gorm.io/gorm"
)

type Message = models.Message

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	res := r.db.Model(Message{}).Where("id = ?", id).Update("text", message.Text)
	if err := res.Error; err != nil {
		return Message{}, err
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	res := r.db.Delete(&Message{}, id)

	if err := res.Error; err != nil {
		return err
	}
	return nil
}
