package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/images"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

type repository struct {
	db *sqlx.DB
	s3 images.Repository
}

func NewChatMemoryRepository(db *sqlx.DB, s3 images.Repository) chat.Repository {
	return &repository{db: db, s3: s3}
}

func (r repository) DeleteChatMembers(ctx context.Context, chatID uint64) error {
	rows, err := r.db.Query("DELETE FROM chat_members WHERE id_chat=$1", chatID)
	defer rows.Close()

	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err
}

func (r repository) UpdateChatAvatar(ctx context.Context, url string, chatID uint64) (model.Chat, error) {
	result, err := r.db.Exec("UPDATE chat SET avatar=$1 WHERE id=$2", url, chatID)
	if err != nil {
		return model.Chat{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Chat{}, err
	}

	if rowsAffected == 0 {
		return model.Chat{}, myErrors.ErrChatNotFound
	}

	var chat model.Chat
	err = r.db.Get(&chat, "SELECT * FROM chat WHERE id=$1", chatID)
	if err != nil {
		return model.Chat{}, err
	}

	return chat, nil
}

func (r repository) UpdateChatById(ctx context.Context, title string, chatID uint64) (model.DBChat, error) {
	var chat model.DBChat
	rows, err := r.db.Query(`UPDATE chat SET title=$1 WHERE id=$2`, title, chatID)
	defer rows.Close()

	if err != nil {
		return model.DBChat{}, err
	}
	if rows.Next() {
		err = rows.Scan(&chat.Id, &chat.Title, &chat.Type, &chat.Avatar)
		if err != nil {
			return model.DBChat{}, err
		}
	}

	return chat, nil
}

func (r repository) GetChatsByUserId(ctx context.Context, userID uint64) ([]model.ChatMembers, error) {
	var chat []model.ChatMembers
	rows, err := r.db.Query("SELECT * FROM chat_members WHERE id_member=$1", userID)
	defer rows.Close()

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
	var chat []model.Chat
	rows, err := r.db.Query("SELECT * FROM chat WHERE id=$1", chatID)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Chat{}, myErrors.ErrChatNotFound
		}
		return model.Chat{}, err
	}

	for rows.Next() {
		var chatFromDB model.Chat
		rows.Scan(&chatFromDB.Id, &chatFromDB.MasterID, &chatFromDB.Type, &chatFromDB.Title, &chatFromDB.Avatar)
		if err != nil {
			return model.Chat{}, err
		}

		chat = append(chat, chatFromDB)
	}

	if len(chat) == 0 {
		return model.Chat{}, nil
	}

	return chat[0], nil
}

func (r repository) CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error) {
	rows, err := r.db.Query(`INSERT INTO chat (master_id, type, avatar, title)  VALUES($1, $2, $3, $4) RETURNING id`,
		chat.MasterID, chat.Type, "", chat.Title)
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
		err = r.AddUserInChatDB(context.TODO(), chat.Id, members.Id)
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
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrChatNotFound
	}

	rows, err = r.db.Query("DELETE FROM message WHERE id_chat=$1", chatID)
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrMessageNotFound
	}

	rows, err = r.db.Query("DELETE FROM chat WHERE id=$1", chatID)
	defer rows.Close()
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
	var groups []model.Chat
	err := r.db.Select(&groups, `
		SELECT id, type, avatar, title 
		FROM chat WHERE type != $1 AND title ILIKE $2 AND 
		EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`,
		configs.Chat, "%"+string+"%", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return groups, nil
}

func (r repository) GetSearchChannels(ctx context.Context, string string, userID uint64) ([]model.Chat, error) {
	var channels []model.Chat
	err := r.db.Select(&channels, `
		SELECT id, type, avatar, title 
		FROM chat WHERE type = $1 AND title ILIKE $2 AND 
		NOT EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`,
		configs.Channel, "%"+string+"%", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return channels, nil
}
