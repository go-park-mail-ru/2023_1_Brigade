package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"regexp"
	"testing"
)

func TestPostgres_CheckExistUserByEmail_True(t *testing.T) {
	var ctx echo.Context
	inputEmail := "CorrectEmail@mail.ru"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"email"}).
		AddRow(inputEmail)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=?`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistEmail(ctx, inputEmail)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByEmail_False(t *testing.T) {
	var ctx echo.Context
	inputEmail := "CorrectEmail@mail.ru"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"email"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=?`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistEmail(ctx, inputEmail)
	require.Error(t, err, myErrors.ErrUserNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByUsername_True(t *testing.T) {
	var ctx echo.Context
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"username"}).
		AddRow(username)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE username=?`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistUsername(ctx, username)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserByUsername_False(t *testing.T) {
	var ctx echo.Context
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"username"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE username=?`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckExistUsername(ctx, username)
	require.Error(t, err, myErrors.ErrUserNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckCorrectPassword_True(t *testing.T) {
	var ctx echo.Context
	email := "marcussss@gmail.com"
	password := "baumanka"
	user := model.User{
		Email:    email,
		Password: password,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"email", "password"}).
		AddRow(email, password)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=? AND password=?`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckCorrectPassword(ctx, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckCorrectPassword_False(t *testing.T) {
	var ctx echo.Context
	email := "marcussss@gmail.com"
	password := "baumanka"
	user := model.User{
		Email:    email,
		Password: password,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	rowMain := sqlmock.NewRows([]string{"email", "password"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=? AND password=?`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthUserMemoryRepository(dbx)

	err = repo.CheckCorrectPassword(ctx, user)
	require.Error(t, err, myErrors.ErrIncorrectPassword)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
