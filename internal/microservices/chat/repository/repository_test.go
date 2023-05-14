package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	"project/internal/model"
	"regexp"
	"testing"
)

func TestPostgres_DeleteChatMembers_OK(t *testing.T) {
	chatID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id_chat"}).
		AddRow(chatID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`DELETE FROM chat_members WHERE id_chat=$1`)).
		WithArgs(chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	err = repo.DeleteChatMembers(context.TODO(), chatID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetChatMembersByChatId_OK(t *testing.T) {
	chatID := uint64(1)
	memberID := uint64(1)
	expectedChatMembers := []model.ChatMembers{
		{
			ChatId:   chatID,
			MemberId: memberID,
		},
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id_chat", "id_member"}).
		AddRow(chatID, memberID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM chat_members WHERE id_chat=$1`)).
		WithArgs(chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	chatMembers, err := repo.GetChatMembersByChatId(context.TODO(), chatID)
	require.Equal(t, expectedChatMembers, chatMembers)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetChatById_OK(t *testing.T) {
	chatID := uint64(0)
	expectedChat := model.Chat{
		Id:     chatID,
		Type:   config.Chat,
		Avatar: "",
		Title:  "",
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "type", "avatar", "title"}).
		AddRow(chatID, config.Chat, "", "")

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM chat WHERE id=$1`)).
		WithArgs(chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	chat, err := repo.GetChatById(context.TODO(), chatID)
	require.Equal(t, expectedChat, chat)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_AddUserInChatDB_OK(t *testing.T) {
	chatID := uint64(1)
	memberID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id_chat", "id_member"}).
		AddRow(chatID, memberID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)`)).
		WithArgs(chatID, memberID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	err = repo.AddUserInChatDB(context.TODO(), chatID, memberID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetChatsByUserId_OK(t *testing.T) {
	chatID := uint64(1)
	userID := uint64(1)
	expectedChats := []model.ChatMembers{
		{
			ChatId:   chatID,
			MemberId: userID,
		},
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id_chat", "id_member"}).
		AddRow(chatID, userID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM chat_members WHERE id_member=$1`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	chats, err := repo.GetChatsByUserId(context.TODO(), userID)
	require.Equal(t, expectedChats, chats)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_UpdateChatById_OK(t *testing.T) {
	title := ""
	chatID := uint64(1)
	expectedChat := model.DBChat{
		Id:     1,
		Type:   config.Chat,
		Title:  title,
		Avatar: "",
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "title", "type", "avatar"}).
		AddRow(1, title, config.Chat, "")

	mock.
		ExpectQuery(regexp.QuoteMeta(`UPDATE chat SET title=$1 WHERE id=$2`)).
		WithArgs(title, chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	chat, err := repo.UpdateChatById(context.TODO(), title, chatID)
	require.NoError(t, err)

	expectedChat.Id = chat.Id
	expectedChat.Type = chat.Type
	expectedChat.Title = chat.Title
	expectedChat.Avatar = chat.Avatar

	require.Equal(t, expectedChat, chat)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
