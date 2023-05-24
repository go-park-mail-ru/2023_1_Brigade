package redis

import (
	"context"
	"github.com/go-redis/redismock/v9"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"testing"
)

func TestRedis_GetSession_OK(t *testing.T) {
	userID := "1"
	cookie := uuid.New().String()
	expectedSession := model.Session{
		UserId: 1,
		Cookie: cookie,
	}

	mockedClient, mock := redismock.NewClientMock()
	repo := NewAuthSessionMemoryRepository(mockedClient)

	mock.ExpectGet(cookie).SetVal(userID)

	session, err := repo.GetSessionByCookie(context.TODO(), cookie)

	require.NoError(t, err)
	require.Equal(t, expectedSession, session)
}

func TestRedis_GetSession_NotFound(t *testing.T) {
	cookie := uuid.New().String()
	mockedClient, mock := redismock.NewClientMock()
	repo := NewAuthSessionMemoryRepository(mockedClient)

	mock.ExpectGet(cookie).SetErr(redis.Nil)

	_, err := repo.GetSessionByCookie(context.TODO(), cookie)

	require.Error(t, err, myErrors.ErrSessionNotFound)
}

func TestRedis_DeleteSession_OK(t *testing.T) {
	cookie := uuid.New().String()
	mockedClient, mock := redismock.NewClientMock()
	repo := NewAuthSessionMemoryRepository(mockedClient)

	mock.ExpectDel(cookie).SetVal(1)

	err := repo.DeleteSession(context.TODO(), cookie)

	require.NoError(t, err)
}

func TestRedis_DeleteSession_NotFound(t *testing.T) {
	cookie := uuid.New().String()
	mockedClient, mock := redismock.NewClientMock()
	repo := NewAuthSessionMemoryRepository(mockedClient)

	mock.ExpectDel(cookie).SetErr(redis.Nil)

	err := repo.DeleteSession(context.TODO(), cookie)

	require.Error(t, err, myErrors.ErrSessionNotFound)
}

func TestRedis_CreateSession_OK(t *testing.T) {
	session := model.Session{
		UserId: 1,
		Cookie: uuid.New().String(),
	}

	mockedClient, mock := redismock.NewClientMock()
	repo := NewAuthSessionMemoryRepository(mockedClient)

	mock.ExpectSet(session.Cookie, session.UserId, 0).SetVal(session.Cookie)

	err := repo.CreateSession(context.TODO(), session)

	require.NoError(t, err)
}
