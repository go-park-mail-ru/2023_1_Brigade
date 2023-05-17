package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	"project/internal/model"
	"regexp"
	"testing"
	"time"
)

func TestPostgres_EditMessageById_OK(t *testing.T) {
	chatID := uint64(1)
	authorID := uint64(1)
	expectedMessage := model.Message{
		Id:        "1",
		Body:      "Hello world!",
		AuthorId:  authorID,
		ChatId:    chatID,
		CreatedAt: "",
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "body", "author_id", "id_chat", "created_at"}).
		AddRow("0", "Bye world!", 1, 1, "")

	mock.
		ExpectQuery(regexp.QuoteMeta(`UPDATE message SET body = $1 WHERE id = $2 RETURNING *`)).
		WithArgs(expectedMessage.Body, expectedMessage.Id).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	message, err := repo.EditMessageById(context.TODO(), model.ProducerMessage{
		Id:         expectedMessage.Id,
		Type:       config.Edit,
		Body:       expectedMessage.Body,
		AuthorId:   expectedMessage.AuthorId,
		ChatID:     expectedMessage.ChatId,
		ReceiverID: 1,
		CreatedAt:  expectedMessage.CreatedAt,
	})
	require.NoError(t, err)

	expectedMessage.Id = message.Id
	expectedMessage.Body = message.Body
	expectedMessage.AuthorId = message.AuthorId
	expectedMessage.ChatId = message.ChatId
	expectedMessage.CreatedAt = message.CreatedAt

	require.Equal(t, expectedMessage, message)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteMessageById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	messageID := uuid.New().String()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM message WHERE id=$1`)).
		WithArgs(messageID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM chat_messages WHERE id_message=$1`)).
		WithArgs(messageID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.DeleteMessageById(context.Background(), messageID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetMessageByID_OK(t *testing.T) {
	expectedMessage := model.Message{
		Id:        uuid.New().String(),
		Body:      "Hello world!",
		AuthorId:  1,
		ChatId:    1,
		CreatedAt: time.Now().String(),
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "body", "author_id", "id_chat", "created_at"}).
		AddRow(expectedMessage.Id, expectedMessage.Body, expectedMessage.AuthorId, expectedMessage.ChatId, expectedMessage.CreatedAt)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM message WHERE id=$1`)).
		WithArgs(expectedMessage.Id).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	message, err := repo.GetMessageById(context.TODO(), expectedMessage.Id)
	require.NoError(t, err)
	require.Equal(t, expectedMessage, message)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetChatMessages_OK(t *testing.T) {
	chatID := uint64(1)
	messageID := uuid.New().String()
	expectedChatMessages := []model.ChatMessages{
		{
			ChatId:    chatID,
			MessageId: messageID,
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

	row := sqlmock.NewRows([]string{"id_chat", "id_message"}).
		AddRow(chatID, messageID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM chat_messages WHERE id_chat=$1`)).
		WithArgs(chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	chatMessages, err := repo.GetChatMessages(context.TODO(), chatID)
	require.NoError(t, err)
	require.Equal(t, expectedChatMessages, chatMessages)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_InsertMessageInDB_OK(t *testing.T) {
	message := model.Message{
		Id:        uuid.New().String(),
		Body:      "Hello world!",
		AuthorId:  1,
		ChatId:    1,
		CreatedAt: time.Now().String(),
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO message (id, body, id_chat, author_id, created_at) VALUES (?, ?, ?, ?, ?)")).
		WithArgs(message.Id, message.Body, message.ChatId, message.AuthorId, message.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO chat_messages (id_chat, id_message) VALUES (?, ?)")).
		WithArgs(message.ChatId, message.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.InsertMessageInDB(context.Background(), message)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetLastChatMessage_OK(t *testing.T) {
	expectedMessage := model.Message{
		Id:        uuid.New().String(),
		Body:      "Hello world!",
		AuthorId:  1,
		ChatId:    1,
		CreatedAt: time.Now().String(),
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "body", "author_id", "id_chat", "created_at"}).
		AddRow(expectedMessage.Id, expectedMessage.Body, expectedMessage.AuthorId, expectedMessage.ChatId, expectedMessage.CreatedAt)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM message WHERE id_chat = $1 AND created_at = (SELECT MAX(created_at) FROM message WHERE id_chat = $1)`)).
		WithArgs(expectedMessage.ChatId).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	message, err := repo.GetLastChatMessage(context.TODO(), expectedMessage.ChatId)
	require.NoError(t, err)
	require.Equal(t, expectedMessage, message)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetSearchMessages_OK(t *testing.T) {
	userID := uint64(1)
	searchString := "abc"
	expectedMessages := []model.Message{
		{
			Id:        uuid.New().String(),
			Body:      "SaBcR!",
			AuthorId:  1,
			ChatId:    1,
			CreatedAt: time.Now().String(),
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

	row := sqlmock.NewRows([]string{"id", "body", "author_id", "id_chat", "created_at"}).
		AddRow(expectedMessages[0].Id, expectedMessages[0].Body, expectedMessages[0].AuthorId, expectedMessages[0].ChatId, expectedMessages[0].CreatedAt)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT message.*
		FROM message
		JOIN chat_messages ON message.id = chat_messages.id_message
		JOIN chat_members ON chat_messages.id_chat = chat_members.id_chat
		WHERE message.body ILIKE $1 AND chat_members.id_member = $2;`)).
		WithArgs("%"+searchString+"%", userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewMessagesMemoryRepository(dbx)

	searchMessages, err := repo.GetSearchMessages(context.TODO(), userID, searchString)
	require.NoError(t, err)
	require.Equal(t, expectedMessages, searchMessages)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
