package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/messages"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
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

func (r *repository) GetChatById(chatID uint64) ([]model.ChatMembers, error) {
	var chat []model.ChatMembers
	err := r.db.Select(&chat, "SELECT * FROM chat_members WHERE id_chat=$1", chatID)

	if errors.Is(err, sql.ErrNoRows) {
		return chat, myErrors.ErrChatNotFound
	}

	return chat, err
}

func (r *repository) InsertMessageReceiveInDB(message model.ProducerMessage) error {
	return nil
}
