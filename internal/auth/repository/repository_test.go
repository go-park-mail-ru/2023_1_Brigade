package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"email"}).
		AddRow(inputEmail)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=?`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isExist, err := repo.CheckExistEmail(ctx, inputEmail)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.True(t, isExist)
}

func TestPostgres_CheckExistUserByEmail_False(t *testing.T) {
	var ctx echo.Context
	inputEmail := "CorrectEmail@mail.ru"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"email"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=?`)).
		WithArgs(inputEmail).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isExist, err := repo.CheckExistEmail(ctx, inputEmail)
	require.Error(t, err, myErrors.ErrUserNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.False(t, isExist)
}

func TestPostgres_CheckExistUserByUsername_True(t *testing.T) {
	var ctx echo.Context
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"username"}).
		AddRow(username)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE username=?`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isExist, err := repo.CheckExistUsername(ctx, username)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.True(t, isExist)
}

func TestPostgres_CheckExistUserByUsername_False(t *testing.T) {
	var ctx echo.Context
	username := "marcussss"

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"username"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE username=?`)).
		WithArgs(username).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isExist, err := repo.CheckExistUsername(ctx, username)
	require.Error(t, err, myErrors.ErrUserNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.False(t, isExist)
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
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"email", "password"}).
		AddRow(email, password)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=? AND password=?`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isCorrect, err := repo.CheckCorrectPassword(ctx, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.True(t, isCorrect)
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
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"email", "password"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=? AND password=?`)).
		WithArgs(email, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	isCorrect, err := repo.CheckCorrectPassword(ctx, user)
	require.Error(t, err, myErrors.ErrIncorrectPassword)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

	require.False(t, isCorrect)
}

func TestPostgres_GetSessionByCookie_True(t *testing.T) {
	var ctx echo.Context
	cookie := uuid.New().String()

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"cookie"}).
		AddRow(cookie)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM session WHERE cookie=$1`)).
		WithArgs(cookie).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	_, err = repo.GetSessionByCookie(ctx, cookie)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetSessionByCookie_False(t *testing.T) {
	var ctx echo.Context
	cookie := uuid.New().String()

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"cookie"})

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM session WHERE cookie=$1`)).
		WithArgs(cookie).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	_, err = repo.GetSessionByCookie(ctx, cookie)
	require.Error(t, err, myErrors.ErrSessionNotFound)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CreateUser_True(t *testing.T) {
	var ctx echo.Context
	username := "marcussss"
	email := "marcussss@gmail.com"
	status := "my status"
	password := "baumanka"
	user := model.User{
		Username: username,
		Email:    email,
		Status:   status,
		Password: password,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"username", "email", "status", "password"}).
		AddRow(username, email, status, password)

	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO profile (username, email, status, password)`)).
		WithArgs(username, email, status, password).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	_, err = repo.CreateUser(ctx, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_DeleteSession_True(t *testing.T) {
	var ctx echo.Context
	cookie := uuid.New().String()

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	rowMain := sqlmock.NewRows([]string{"cookie"}).
		AddRow(cookie)

	mock.
		ExpectQuery(regexp.QuoteMeta(`DELETE FROM session WHERE cookie=$1`)).
		WithArgs(cookie).
		WillReturnRows(rowMain)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthMemoryRepository(dbx)

	err = repo.DeleteSession(ctx, cookie)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
