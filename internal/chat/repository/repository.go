package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/chat"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func NewChatMemoryRepository(db *sqlx.DB) chat.Repository {
	return &repository{db: db}
}

func (r repository) GetChatById(ctx context.Context, chatID uint64) (model.Chat, error) {
	var chat model.Chat
	err := r.db.Get(&chat, "SELECT * FROM chat WHERE id=$1", chatID)

	if errors.Is(err, sql.ErrNoRows) {
		return chat, myErrors.ErrChatNotFound
	}

	return chat, err
}

func (r repository) CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error) {
	rows, err := r.db.NamedQuery("INSERT INTO chat (type, name, created_at, members, masters) "+
		"VALUES (:type, :name, :created_at, :members, :masters) RETURNING id", chat)

	if err != nil {
		return model.Chat{}, err
	}
	if rows.Next() {
		err = rows.Scan(&chat.Id)
		if err != nil {
			return model.Chat{}, err
		}
	}

	return chat, nil
}

func (r repository) DeleteChatById(ctx context.Context, chatID uint64) error {
	_, err := r.db.Query("DELETE FROM chat WHERE id=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	return err
}

func (r repository) AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error {
	_, err := r.db.Query("INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)", chatID, memberID)
	if err != nil {
		return err
	}

	return nil
}
