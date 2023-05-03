package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	myErrors "project/internal/pkg/errors"
	"regexp"
	"testing"
)

//func TestPostgres_CreateUser_OK(t *testing.T) {
//	expectedUser := model.AuthorizedUser{
//		Id:       1,
//		Avatar:   "",
//		Username: "marcussss",
//		Nickname: "marcussss",
//		Email:    "marcussss@mail.ru",
//		Password: "12345678",
//	}
//
//	db, mock, err := sqlmock.New()
//	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	defer func() {
//		err := db.Close()
//		if err != nil {
//			log.Error(err)
//		}
//	}()
//
//	row := sqlmock.NewRows([]string{"id", "avatar", "username", "nickname", "email", "status", "password"}).
//		AddRow(expectedUser.Id, expectedUser.Avatar, expectedUser.Username, expectedUser.Nickname, expectedUser.Email, expectedUser.Status, expectedUser.Password)
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`INSERT INTO profile (avatar, username, nickname, email, status, password) `+
//			`VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)).
//		WithArgs(expectedUser.Avatar, expectedUser.Username, expectedUser.Nickname, expectedUser.Email, expectedUser.Status, expectedUser.Password).
//		WillReturnRows(row)
//
//	dbx := sqlx.NewDb(db, "sqlmock")
//	repo := NewAuthUserMemoryRepository(dbx)
//
//	user, err := repo.CreateUser(context.TODO(), expectedUser)
//	require.Equal(t, expectedUser, user)
//	require.NoError(t, err)
//
//	err = mock.ExpectationsWereMet()
//	require.NoError(t, err)
//}

func TestPostgres_CheckExistUserByEmail_True(t *testing.T) {
	inputEmail := "CorrectEmail@mail.ru"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(true)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistEmail(context.Background(), inputEmail)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByEmail_False(t *testing.T) {
	inputEmail := "IncorrectEmail@mail.ru"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistEmail(context.Background(), inputEmail)
	require.Error(t, err, myErrors.ErrEmailNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByUsername_True(t *testing.T) {
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(true)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistUsername(context.Background(), username)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByUsername_False(t *testing.T) {
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE username=$1)`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistUsername(context.Background(), username)
	require.Error(t, err, myErrors.ErrUsernameNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckCorrectPassword_True(t *testing.T) {
	email := "marcussss@gmail.com"
	password := "baumanka"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(true)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckCorrectPassword(context.Background(), email, password)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckCorrectPassword_False(t *testing.T) {
	email := "marcussss@gmail.com"
	password := "baumanka"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1 AND password=$2)`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckCorrectPassword(context.Background(), email, password)
	require.Error(t, err, myErrors.ErrIncorrectPassword)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
