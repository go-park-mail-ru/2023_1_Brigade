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

func (r repository) GetMessageById(ctx context.Context, messageID uint64) (model.Message, error) {
	var message model.Message
	err := r.db.Get(&message, "SELECT * FROM message WHERE id=$1", messageID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, myErrors.ErrMessageNotFound
	}

	return message, err
}

func (r repository) GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error) {
	var chatMessages []model.ChatMessages
	rows, err := r.db.Query("SELECT * FROM chat_messsages WHERE id_chat", chatID)

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
	return message, nil
}

func (r repository) GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error) {
	var lastMessage model.Message
	err := r.db.Get(&lastMessage, `SELECT * FROM message WHERE id_chat=$1`, chatID)

	if errors.Is(err, sql.ErrNoRows) {
		return model.Message{}, nil
	}

	return lastMessage, err
}

//func (r repository) MarkMessageReading(ctx context.Context, messageID uint64) error {
//	return nil
//}

//func (r repository) GetChatById(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
//var chat []model.ChatMembers
//err := r.db.Select(&chat, "SELECT * FROM chat_members WHERE id_chat=$1", chatID)
//
//if errors.Is(err, sql.ErrNoRows) {
//	return []model.ChatMembers{}, myErrors.ErrChatNotFound
//}
//
//return chat, err
//}

//func (r repository) InsertMessageReceiveInDB(ctx context.Context, message model.ProducerMessage) error {
//	return nil
//}
