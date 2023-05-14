package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"project/internal/microservices/messages"
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
	err := r.db.GetContext(ctx, &message, "UPDATE message SET body = $1 WHERE id = $ RETURNING *", producerMessage.Body, producerMessage.Id)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}

func (r repository) DeleteMessageById(ctx context.Context, messageID string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM message WHERE id=$1", messageID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat_messages WHERE id_message=$1", messageID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetMessageById(ctx context.Context, messageID string) (model.Message, error) {
	var message model.Message
	err := r.db.GetContext(ctx, &message, "SELECT * FROM message WHERE id=$1", messageID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, myErrors.ErrMessageNotFound
		}

		return model.Message{}, err
	}

	return message, nil
}

func (r repository) GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error) {
	var chatMessages []model.ChatMessages
	err := r.db.SelectContext(ctx, &chatMessages, "SELECT * FROM chat_messages WHERE id_chat=$1", chatID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrMembersNotFound
		}

		return nil, err
	}

	return chatMessages, nil
}

func (r repository) InsertMessageInDB(ctx context.Context, message model.Message) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = r.db.NamedExecContext(ctx, `INSERT INTO message (id, body, id_chat, author_id, created_at) `+
		`VALUES (:id, :body, :id_chat, :author_id, :created_at)`, message)
	log.Info(err)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.NamedExecContext(ctx, "INSERT INTO chat_messages (id_chat, id_message) VALUES (:id_chat, :id_message)", model.ChatMessages{
		ChatId:    message.ChatId,
		MessageId: message.Id,
	})
	log.Info(err)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	log.Info(err)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error) {
	var lastMessage model.Message
	err := r.db.Get(&lastMessage, `SELECT * FROM message WHERE id_chat = $1 AND created_at = (SELECT MAX(created_at) FROM message WHERE id_chat = $1)`, chatID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, nil
		}

		return model.Message{}, err
	}

	return lastMessage, nil
}

func (r repository) GetSearchMessages(ctx context.Context, userID uint64, string string) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(&messages, `
		SELECT message.*
		FROM message
		JOIN chat_messages ON message.id = chat_messages.id_message
		JOIN chat_members ON chat_messages.id_chat = chat_members.id_chat
		WHERE message.body ILIKE $1 AND chat_members.id_member = $2;`,
		"%"+string+"%", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return messages, nil
}
