package postgres

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"project/internal/model"
	"regexp"
	"testing"
)

func TestPostgres_GetSession_OK(t *testing.T) {
	cookie := uuid.New().String()
	expectedSession := model.Session{
		UserId: 1,
		Cookie: cookie,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	row := sqlmock.NewRows([]string{"profile_id", "cookie"}).
		AddRow(expectedSession.UserId, expectedSession.Cookie)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM session WHERE cookie = $1`)).
		WithArgs(expectedSession.Cookie).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthSessionMemoryRepository(dbx)

	session, err := repo.GetSessionByCookie(context.TODO(), cookie)
	require.Equal(t, expectedSession, session)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_CreateSession_OK(t *testing.T) {
	cookie := uuid.New().String()
	session := model.Session{
		UserId: 1,
		Cookie: cookie,
	}

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO session (cookie, profile_id) VALUES ($1, $2)`)).
		WithArgs(session.Cookie, session.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthSessionMemoryRepository(dbx)

	err = repo.CreateSession(context.TODO(), session)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostgres_DeleteSession_OK(t *testing.T) {
	cookie := uuid.New().String()

	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	//row := sqlmock.NewRows([]string{"session", "cookie"}).
	//	AddRow(userID, cookie)
	//
	//mock.
	//	ExpectQuery(regexp.QuoteMeta(`DELETE FROM session WHERE cookie=$1`)).
	//	WithArgs(cookie).
	//	WillReturnRows(row)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM session WHERE cookie=$1")).
		WithArgs(cookie).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewAuthSessionMemoryRepository(dbx)

	err = repo.DeleteSession(context.TODO(), cookie)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
