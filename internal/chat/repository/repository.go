package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
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

func (r repository) GetMembersByChatId(ctx context.Context, chatID uint64) ([]model.User, error) {
	var members []model.User
	rows, err := r.db.Query("SELECT * FROM chat_members WHERE id_chat=$1", chatID)
	//
	if err != nil {
		// TODO
		if errors.Is(err, sql.ErrNoRows) {
			log.Error(err)
			return nil, nil
			//return nil, myErrors.ErrChatNotFound
		}
		return nil, err
	}

	for rows.Next() {
		var member model.User
		err := rows.Scan(&member.Id, &member.Username, &member.Nickname, &member.Email, &member.Status, &member.Avatar)
		if err != nil {
			log.Error(err)
		}
		members = append(members, member)
	}

	return members, err
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

	log.Warn(chat)
	for _, members := range chat.Members {
		err = r.AddUserInChatDB(context.Background(), chat.Id, members.Id)
		if err != nil {
			log.Error(err)
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

	//_, err = r.db.Query("INSERT INTO users_chats (id_user, id_chat) VALUES ($1, $2)", memberID, chatID)
	//if err != nil {
	//	return err
	//}

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
			log.Error(err)
		}
		chat = append(chat, memberChat)
	}

	return chat, err
}
