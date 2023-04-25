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

func (r repository) DeleteChatMembers(ctx context.Context, chatID uint64) error {
	_, err := r.db.Query("DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r repository) UpdateChatById(ctx context.Context, title string, chatID uint64) (model.DBChat, error) {
	var chat model.DBChat
	result, err := r.db.Exec("UPDATE chat SET title=$1 WHERE id=$2", title, chatID)
	if err != nil {
		return model.DBChat{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.DBChat{}, err
	}

	if rowsAffected == 0 {
		return model.DBChat{}, err
	}

	err = r.db.Get(&chat, "SELECT * FROM chat WHERE id=$1", chatID)
	if err != nil {
		return model.DBChat{}, err
	}

	return chat, nil
}

func (r repository) GetChatMembersByChatId(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
	var chatMembers []model.ChatMembers
	rows, err := r.db.Query("SELECT * FROM chat_members WHERE id_chat=$1", chatID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrMembersNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var chatMember model.ChatMembers
		err := rows.Scan(&chatMember.ChatId, &chatMember.MemberId)
		if err != nil {
			return nil, err
		}

		chatMembers = append(chatMembers, chatMember)
	}

	return chatMembers, err
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
	rows, err := r.db.NamedQuery("INSERT INTO chat (type, title, avatar) "+
		"VALUES (:type, :title, :avatar) RETURNING id", chat)

	if err != nil {
		return model.Chat{}, err
	}
	if rows.Next() {
		err = rows.Scan(&chat.Id)
		if err != nil {
			return model.Chat{}, err
		}
	}

	for _, members := range chat.Members {
		err = r.AddUserInChatDB(context.Background(), chat.Id, members.Id)
		if err != nil {
			return model.Chat{}, err
		}
	}

	return chat, nil
}

func (r repository) DeleteChatById(ctx context.Context, chatID uint64) error {
	_, err := r.db.Query("DELETE FROM chat_messages WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	_, err = r.db.Query("DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	_, err = r.db.Query("DELETE FROM message WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrMessageNotFound
	}

	_, err = r.db.Query("DELETE FROM chat WHERE id=$1", chatID)
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

func (r repository) GetChatsByUserId(ctx context.Context, userID uint64) ([]model.ChatMembers, error) {
	var chat []model.ChatMembers
	rows, err := r.db.Query("SELECT * FROM chat_members WHERE id_member=$1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrChatNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var memberChat model.ChatMembers
		err := rows.Scan(&memberChat.ChatId, &memberChat.MemberId)
		if err != nil {
			return nil, err
		}
		chat = append(chat, memberChat)
	}

	return chat, err
}
