package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/chat"
	"project/internal/configs"
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
	rows, err := r.db.Query("DELETE FROM chat_members WHERE id_chat=$1", chatID)
	defer rows.Close()

	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r repository) UpdateChatById(ctx context.Context, title string, chatID uint64) (model.DBChat, error) {
	var chat model.DBChat
	rows, err := r.db.Query(`UPDATE chat SET title=$1 WHERE id=$2`, title, chatID)
	defer rows.Close()

	if err != nil {
		return model.DBChat{}, err
	}
	if rows.Next() {
		err = rows.Scan(&chat)
		if err != nil {
			return model.DBChat{}, err
		}
	}

	return chat, nil
}

func (r repository) GetChatMembersByChatId(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
	var chatMembers []model.ChatMembers
	rows, err := r.db.Query("SELECT * FROM chat_members WHERE id_chat=$1", chatID)
	defer rows.Close()

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
	defer rows.Close()

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
	rows, err := r.db.Query("DELETE FROM chat_messages WHERE id_chat=$1", chatID)
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	rows, err = r.db.Query("DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	rows, err = r.db.Query("DELETE FROM message WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrMessageNotFound
	}

	rows, err = r.db.Query("DELETE FROM chat WHERE id=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	return err
}

func (r repository) AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error {
	rows, err := r.db.Query("INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)", chatID, memberID)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetSearchChats(ctx context.Context, userID uint64, string string) ([]model.Chat, error) {
	return nil, nil
	//var chat model.Chat
	//err := r.db.Get(&chat, "SELECT chat.id FROM chat JOIN chat_members ON chat.id = chat_members.id_chat WHERE chat_members.id_member = $1 AND chat.title LIKE $2;", userID, "%"+string+"%")
	//
	//if errors.Is(err, sql.ErrNoRows) {
	//return chat, myErrors.ErrChatNotFound
	//}

	//return chat, err
}

func (r repository) GetSearchChannels(ctx context.Context, string string) ([]model.Chat, error) {
	var channels []model.Chat
	err := r.db.Select(&channels, `SELECT * FROM chat WHERE title LIKE $1 AND type=$2`, "%"+string+"%", configs.Channel)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return channels, nil
}
