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

func TestPostgres_CreateTechnogrammChat_OK(t *testing.T) {
	user := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Hello world!",
		Password: "12345678",
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
	repo := NewChatMemoryRepository(dbx)

	row := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO chat (master_id, type, avatar, title) 
    VALUES (0, 0, 'https://brigade_chat_avatars.hb.bizmrg.com/logo.png', 'Technogramm') RETURNING id;`)).
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO message (id, type, body, id_chat, author_id, created_at)
    VALUES ($1, $2, 'Привет, это технограмм!', (SELECT id FROM chat WHERE id = $3), (SELECT id FROM profile WHERE id = $4), '0001-01-01 00:00:00+00');`)).
		WithArgs(sqlmock.AnyArg(), config.NotSticker, 1, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO chat_messages (id_chat, id_message)
    VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM message WHERE id = $2));`)).
		WithArgs(1, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO chat_members (id_chat, id_member)
   VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = $2));`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO chat_members (id_chat, id_member)
   VALUES ((SELECT id FROM chat WHERE id = $1), (SELECT id FROM profile WHERE id = 0));`)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.CreateTechnogrammChat(context.TODO(), user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

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

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM chat_members WHERE id_chat=$1")).
		WithArgs(chatID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	err = repo.DeleteChatMembers(context.TODO(), chatID)
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

func TestPostgres_CreateChat_OK(t *testing.T) {
	chat := model.Chat{
		MasterID: 1337,
		Type:     config.Chat,
		Title:    "Sergey",
		Avatar:   "",
		Members: []model.User{
			{
				Id: 1337,
			},
			{
				Id: 1338,
			},
		},
	}
	expectedChat := model.Chat{
		Id:       1,
		MasterID: 1337,
		Type:     config.Chat,
		Title:    "Sergey",
		Avatar:   "",
		Members: []model.User{
			{
				Id: 1337,
			},
			{
				Id: 1338,
			},
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

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	row := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO chat (master_id, type, avatar, title)  VALUES($1, $2, $3, $4) RETURNING id`)).
		WithArgs(chat.MasterID, chat.Type, chat.Avatar, chat.Title).
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)")).
		WithArgs(1, 1337).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)")).
		WithArgs(1, 1338).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	returnedChat, err := repo.CreateChat(context.TODO(), chat)
	require.NoError(t, err)
	require.Equal(t, expectedChat, returnedChat)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_DeleteChatById_OK(t *testing.T) {
	chatID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM chat_messages WHERE id_chat=$1")).
		WithArgs(chatID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM chat_members WHERE id_chat=$1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM message WHERE id_chat=$1")).
		WithArgs(chatID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM chat WHERE id=$1")).
		WithArgs(chatID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.DeleteChatById(context.TODO(), chatID)
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

	mock.
		ExpectExec(regexp.QuoteMeta(`INSERT INTO chat_members (id_chat, id_member) VALUES ($1, $2)`)).
		WithArgs(chatID, memberID).
		WillReturnResult(sqlmock.NewResult(1, 1))

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

func TestPostgres_UpdateChatAvatar_OK(t *testing.T) {
	chatID := uint64(1)
	url := "vk.com"
	chatInDB := model.DBChat{
		Id:       1,
		MasterID: 1337,
		Type:     config.Group,
		Title:    "chat",
		Avatar:   url,
	}

	expectedChat := model.Chat{
		Id:       1,
		MasterID: 1337,
		Type:     config.Group,
		Title:    "chat",
		Avatar:   url,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "master_id", "type", "title", "avatar"}).
		AddRow(chatInDB.Id, chatInDB.MasterID, chatInDB.Type, chatInDB.Title, chatInDB.Avatar)

	mock.
		ExpectQuery(regexp.QuoteMeta(`UPDATE chat SET avatar=$1 WHERE id=$2 RETURNING *`)).
		WithArgs(url, chatID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	chat, err := repo.UpdateChatAvatar(context.TODO(), url, chatID)
	require.NoError(t, err)
	require.Equal(t, expectedChat, chat)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetSearchChats_OK(t *testing.T) {
	userID := uint64(1)
	searchString := "abc"
	expectedChats := []model.Chat{
		{
			Id:       1,
			MasterID: 1337,
			Type:     config.Group,
			Title:    "SaBcR!",
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

	row := sqlmock.NewRows([]string{"id", "master_id", "type", "title"}).
		AddRow(expectedChats[0].Id, expectedChats[0].MasterID, expectedChats[0].Type, expectedChats[0].Title)

	mock.
		ExpectQuery(regexp.QuoteMeta(`
		SELECT id, type, avatar, title 
		FROM chat WHERE type != $1 AND title ILIKE $2 AND 
		EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`)).
		WithArgs(config.Chat, "%"+searchString+"%", userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	searchChats, err := repo.GetSearchChats(context.TODO(), userID, searchString)
	require.NoError(t, err)
	require.Equal(t, expectedChats, searchChats)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetSearchChannels_OK(t *testing.T) {
	userID := uint64(1)
	searchString := "abc"
	expectedChats := []model.Chat{
		{
			Id:       1,
			MasterID: 1337,
			Type:     config.Group,
			Title:    "SaBcR!",
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

	row := sqlmock.NewRows([]string{"id", "master_id", "type", "title"}).
		AddRow(expectedChats[0].Id, expectedChats[0].MasterID, expectedChats[0].Type, expectedChats[0].Title)

	mock.
		ExpectQuery(regexp.QuoteMeta(`
		SELECT id, type, avatar, title 
		FROM chat WHERE type = $1 AND title ILIKE $2 AND 
		NOT EXISTS (SELECT 1 FROM chat_members WHERE id_chat = chat.id AND id_member = $3)`)).
		WithArgs(config.Channel, "%"+searchString+"%", userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewChatMemoryRepository(dbx)

	searchChannels, err := repo.GetSearchChannels(context.TODO(), searchString, userID)
	require.NoError(t, err)
	require.Equal(t, expectedChats, searchChannels)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
