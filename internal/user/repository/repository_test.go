package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"regexp"
	"testing"
)

func TestPostgres_DeleteUserById_OK(t *testing.T) {
	userID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "avatar", "username", "nickname", "email", "status", "password"}).
		AddRow(1, "", "marcussss", "marcussss", "marcussss@mail.ru", "Cool status!", "password")

	mock.
		ExpectQuery(regexp.QuoteMeta(`DELETE FROM profile WHERE id=$1`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	err = repo.DeleteUserById(context.TODO(), userID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetUserById_OK(t *testing.T) {
	userID := uint64(1)
	expectedUser := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    "marcussss@mail.ru",
		Status:   "Cool status!",
		Password: "password",
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "avatar", "username", "nickname", "email", "status", "password"}).
		AddRow(1, "", "marcussss", "marcussss", "marcussss@mail.ru", "Cool status!", "password")

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE id=$1`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	user, err := repo.GetUserById(context.TODO(), userID)
	require.Equal(t, expectedUser, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetUserByEmail_OK(t *testing.T) {
	email := "marcussss@mail.ru"
	expectedUser := model.AuthorizedUser{
		Id:       1,
		Avatar:   "",
		Username: "marcussss",
		Nickname: "marcussss",
		Email:    email,
		Status:   "Cool status!",
		Password: "password",
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id", "avatar", "username", "nickname", "email", "status", "password"}).
		AddRow(1, "", "marcussss", "marcussss", "marcussss@mail.ru", "Cool status!", "password")

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM profile WHERE email=$1`)).
		WithArgs(email).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	user, err := repo.GetUserByEmail(context.TODO(), email)
	require.Equal(t, expectedUser, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_GetUserContacts_OK(t *testing.T) {
	//email := "marcussss@mail.ru"
	//expectedUser := model.AuthorizedUser{
	//	Id:       1,
	//	Avatar:   "",
	//	Username: "marcussss",
	//	Nickname: "marcussss",
	//	Email:    email,
	//	Status:   "Cool status!",
	//	Password: "password",
	//}
	userID := uint64(1)
	contactID := uint64(2)
	expectedContacts := []model.AuthorizedUser{model.AuthorizedUser{}}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"id_user", "id_contact"}).
		AddRow(userID, contactID)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user_contacts WHERE id_user=$1`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	contacts, err := repo.GetUserContacts(context.TODO(), userID)
	require.Equal(t, expectedContacts, contacts)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserById_OK(t *testing.T) {
	userID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"exists"}).
		AddRow(true)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE id=$1)`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	err = repo.CheckExistUserById(context.Background(), userID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CheckExistUserById_False(t *testing.T) {
	userID := uint64(1)

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM profile WHERE id=$1)`)).
		WithArgs(userID).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewUserMemoryRepository(dbx)

	err = repo.CheckExistUserById(context.Background(), userID)
	require.Error(t, myErrors.ErrUserNotFound, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
