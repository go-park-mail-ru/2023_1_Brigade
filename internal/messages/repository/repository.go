package repository

import (
	"github.com/jmoiron/sqlx"
	"project/internal/messages"
	"project/internal/model"
)

type repository struct {
	db *sqlx.DB
}

func NewMessagesMemoryRepository(db *sqlx.DB) messages.Repository {
	return &repository{db: db}
}

func (r *repository) InsertMessageInDB(message model.Message) (model.Message, error) {
	return message, nil
}

func (r *repository) MarkMessageReading(messageID uint64) error {
	return nil
}

func (r *repository) GetChatById(chatID uint64) (model.Chat, error) {
	return model.Chat{Members: []model.User{model.User{}}}, nil
}

func (r *repository) InsertMessageReceiveInDB(message model.ProducerMessage) error {
	return nil
}
