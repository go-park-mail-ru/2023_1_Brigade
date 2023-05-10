package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
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
	_, err := r.db.ExecContext(ctx, "DELETE FROM chat_members WHERE id_chat=$1", chatID)

	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err

}

func (r repository) UpdateChatById(ctx context.Context, title string, chatID uint64) (model.DBChat, error) {
	var chat model.DBChat
	err := r.db.GetContext(ctx, &chat, `UPDATE chat SET title=$1 WHERE id=$2`, title, chatID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.DBChat{}, myErrors.ErrChatNotFound
		}
		return model.DBChat{}, err
	}

	return chat, nil
}

func (r repository) GetChatsByUserId(ctx context.Context, userID uint64) ([]model.ChatMembers, error) {
	var chatMembers []model.ChatMembers
	err := r.db.SelectContext(ctx, &chatMembers, "SELECT * FROM chat_members WHERE id_member=$1", userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return chatMembers, nil

}

func (r repository) GetChatMembersByChatId(ctx context.Context, chatID uint64) ([]model.ChatMembers, error) {
	var chatMembers []model.ChatMembers
	err := r.db.SelectContext(ctx, &chatMembers, "SELECT * FROM chat_members WHERE id_chat=$1", chatID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrMembersNotFound
		}
		return nil, err
	}

	return chatMembers, nil
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
		rows.Scan(&chatFromDB.Id, &chatFromDB.MasterID, &chatFromDB.Type, &chatFromDB.Avatar, &chatFromDB.Title)
		if err != nil {
			return model.Chat{}, err
		}

		chat = append(chat, chatFromDB)
	}
	log.Info(chat)
	if len(chat) == 0 {
		return model.Chat{}, nil
	}

	return chat[0], nil
}

func (r repository) CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Chat{}, err
	}

	var chatDB model.DBChat
	err = r.db.QueryRowContext(ctx, `INSERT INTO chat (master_id, type, avatar, title)  VALUES($1, $2, $3, $4) RETURNING id`,
		chat.MasterID, chat.Type, "", chat.Title).Scan(&chatDB.Id)
	if err != nil {
		tx.Rollback()
		return model.Chat{}, err
	}
	chat.Id = chatDB.Id

	for _, members := range chat.Members {
		err = r.AddUserInChatDB(ctx, chat.Id, members.Id)
		if err != nil {
			tx.Rollback()
			return model.Chat{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return model.Chat{}, err
	}

	return chat, nil
}

func (r repository) DeleteChatById(ctx context.Context, chatID uint64) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat_messages WHERE id_chat=$1", chatID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM message WHERE id_chat=$1", chatID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat WHERE id=$1", chatID)
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

func (r repository) AddUserInChatDB(ctx context.Context, chatID uint64, memberID uint64) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)", chatID, memberID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateChatAvatar(ctx context.Context, url string, chatID uint64) (model.Chat, error) {
	var chat model.DBChat
	err := r.db.GetContext(ctx, &chat, `UPDATE chat SET avatar=$1 WHERE id=$2 RETURNING *`, url, chatID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Chat{}, myErrors.ErrChatNotFound
		}
		return model.Chat{}, err
	}

	return model.Chat{
		Id:       chat.Id,
		MasterID: chat.MasterID,
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   chat.Avatar,
	}, nil
}

func (r repository) GetSearchChats(ctx context.Context, userID uint64, string string) ([]model.Chat, error) {
	var groups []model.Chat
	err := r.db.SelectContext(ctx, &groups, `
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
