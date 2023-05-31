package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"project/internal/config"
	"project/internal/microservices/chat"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func NewChatMemoryRepository(db *sqlx.DB) chat.Repository {
	return &repository{db: db}
}

func (r repository) CreateTechnogrammChat(ctx context.Context, user model.AuthorizedUser) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	var chat model.DBChat
	err = tx.QueryRowContext(ctx, `INSERT INTO chat (master_id, type, description, avatar, title)
   VALUES (0, 0, 'Служебный чат', 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm') RETURNING id;`).Scan(&chat.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	id := uuid.New().String()
	_, err = tx.ExecContext(ctx, `INSERT INTO message (id, type, body, id_chat, author_id, created_at)
   VALUES ($1, $2, $3, (SELECT id FROM chat WHERE id = $4), (SELECT id FROM profile WHERE id = $5), '0001-01-01 00:00:00+00');`, id, config.NotSticker,
		"Привет, это технограмм - бесплатный мессенджер, разработанный студентами технопарка.\n\n"+
			"Здесь ты можешь "+
			"обмениваться текстовыми сообщениями, фото, стикерами, эмодзи, с одним или несколькими пользователями. "+
			"Создавать групповой чат, добавлять и удалять из него пользователей. "+
			"Создавать канал, на него может подписаться пользователь, и при публикации получит уведомление.\n\n"+
			"Желаем тебе удачного общения!\n", chat.Id, 0)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO chat_messages (id_chat, id_message)
   VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM message WHERE id = $2));`, chat.Id, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO chat_members (id_chat, id_member)
   VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = $2));`, chat.Id, user.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO chat_members (id_chat, id_member)
   VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = 0));`, chat.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteChatMembers(ctx context.Context, chatID uint64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if errors.Is(err, sql.ErrNoRows) {
		return myErrors.ErrUserNotFound
	}

	return err

}

func (r repository) UpdateChatById(ctx context.Context, editChat model.EditChat) (model.DBChat, error) {
	var chat model.DBChat
	err := r.db.GetContext(ctx, &chat, `UPDATE chat SET description=$1, avatar=$2, title=$3 WHERE id=$4 RETURNING *`,
		editChat.Description, editChat.Avatar, editChat.Title, editChat.Id)
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
	var chat model.DBChat
	err := r.db.GetContext(ctx, &chat, "SELECT * FROM chat WHERE id=$1", chatID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Chat{}, myErrors.ErrMembersNotFound
		}

		return model.Chat{}, err
	}

	return model.Chat{
		Id:          chat.Id,
		MasterID:    chat.MasterID,
		Type:        chat.Type,
		Description: chat.Description,
		Title:       chat.Title,
		Avatar:      chat.Avatar,
	}, nil
}

func (r repository) CreateChat(ctx context.Context, chat model.Chat) (model.Chat, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Chat{}, err
	}

	var chatDB model.DBChat
	err = r.db.QueryRowContext(ctx, `INSERT INTO chat (master_id, type, description, avatar, title)  VALUES($1, $2, $3, $4, $5) RETURNING id`,
		chat.MasterID, chat.Type, chat.Description, chat.Avatar, chat.Title).Scan(&chatDB.Id)
	if err != nil {
		_ = tx.Rollback()
		return model.Chat{}, err
	}
	chat.Id = chatDB.Id

	for _, members := range chat.Members {
		err = r.AddUserInChatDB(ctx, chat.Id, members.Id)
		if err != nil {
			_ = tx.Rollback()
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
		_ = tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat_members WHERE id_chat=$1", chatID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM message WHERE id_chat=$1", chatID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM chat WHERE id=$1", chatID)
	if err != nil {
		_ = tx.Rollback()
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
		Id:          chat.Id,
		MasterID:    chat.MasterID,
		Type:        chat.Type,
		Description: chat.Description,
		Title:       chat.Title,
		Avatar:      chat.Avatar,
	}, nil
}

func (r repository) GetSearchChats(ctx context.Context, userID uint64, string string) ([]model.Chat, error) {
	var groups []model.Chat
	err := r.db.SelectContext(ctx, &groups, `
		SELECT id, type, description, avatar, title 
		FROM chat WHERE type != $1 AND title ILIKE $2 AND 
		EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`,
		config.Chat, "%"+string+"%", userID)
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
	err := r.db.SelectContext(ctx, &channels, `
		SELECT id, type, description, avatar, title 
		FROM chat WHERE type = $1 AND title ILIKE $2 AND 
		NOT EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`,
		config.Channel, "%"+string+"%", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myErrors.ErrChatNotFound
		}

		return nil, err
	}

	return channels, nil
}
