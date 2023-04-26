package repository

import (
	"context"
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

func (r repository) EditMessageById(ctx context.Context, producerMessage model.ProducerMessage) (model.Message, error) {
	var message model.Message
	err := r.db.Get(&message, "UPDATE message SET body = $1, created_at = $2 WHERE id = $3 RETURNING *", producerMessage.Body, producerMessage.CreatedAt, producerMessage.Id)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}

func (r repository) DeleteMessageById(ctx context.Context, messageID string) error {
	_, err := r.db.Exec("DELETE FROM message WHERE id=$1", messageID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrMessageNotFound
	}

	_, err = r.db.Exec("DELETE FROM chat_messages WHERE id_message=$1", messageID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrMessageNotFound
	}

	return nil
}

func (r repository) GetMessageById(ctx context.Context, messageID string) (model.Message, error) {
	var message model.Message
	err := r.db.Get(&message, "SELECT * FROM message WHERE id=$1", messageID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, myErrors.ErrMessageNotFound
	}

	return message, err
}

func (r repository) GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error) {
	var chatMessages []model.ChatMessages
	rows, err := r.db.Query("SELECT * FROM chat_messages WHERE id_chat=$1", chatID)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrChatNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var chatMessage model.ChatMessages
		err := rows.Scan(&chatMessage.ChatId, &chatMessage.MessageId)
		if err != nil {
			return nil, err
		}

		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages, err
}

func (r repository) InsertMessageInDB(ctx context.Context, message model.Message) (model.Message, error) {
	rows, err := r.db.NamedQuery(`INSERT INTO message (id, body, id_chat, author_id, created_at) `+
		`VALUES (:id, :body, :id_chat, :author_id, :created_at)`, message)
	defer rows.Close()

	if err != nil {
		return model.Message{}, err
	}

	rows, err = r.db.NamedQuery("INSERT INTO chat_messages (id_chat, id_message) VALUES (:id_chat, :id_message)", model.ChatMessages{
		ChatId:    message.ChatId,
		MessageId: message.Id,
	})
	defer rows.Close()

	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}

func (r repository) GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error) {
	var lastMessage model.Message
	err := r.db.Get(&lastMessage, `SELECT * FROM message WHERE id_chat = $1 AND created_at = (SELECT MAX(created_at) FROM message WHERE id_chat = $1)`, chatID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, nil
	}

	return lastMessage, err
}
