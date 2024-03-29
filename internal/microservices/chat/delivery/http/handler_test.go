package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	chatMock "project/internal/microservices/chat/usecase/mocks"
	userMock "project/internal/microservices/user/usecase/mocks"
	"project/internal/model"
	"testing"
)

type testCase struct {
	body   []byte
	status int
	name   string
}

func TestHandlers_CreateChat_OK(t *testing.T) {
	test := testCase{[]byte(`{"title": "chat_title",
								   "members": [0]}`),
		http.StatusCreated,
		"Successfull creating chat"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/chats/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})

	var chat model.CreateChat
	var dbChat model.Chat
	chatUsecase := chatMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewChatHandler(e, chatUsecase, userUsecase)

	err := json.Unmarshal(test.body, &chat)
	require.NoError(t, err)

	chatUsecase.EXPECT().CreateChat(context.TODO(), chat, uint64(1)).Return(dbChat, nil).Times(1)

	err = handler.CreateCurrentUserChatHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_GetChat_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusOK,
		"Successfull getting chat"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/chats/1", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: uint64(1)})
	ctx.SetParamNames("chatID")
	ctx.SetParamValues("1")

	var chat model.Chat
	chatUsecase := chatMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewChatHandler(e, chatUsecase, userUsecase)

	chatUsecase.EXPECT().GetChatById(context.TODO(), uint64(1), uint64(1)).Return(chat, nil).Times(1)

	err := handler.GetChatHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_DeleteChat_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNoContent,
		"Successfull deleting chat"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/chats/1", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{})
	ctx.SetParamNames("chatID")
	ctx.SetParamValues("1")

	chatUsecase := chatMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewChatHandler(e, chatUsecase, userUsecase)

	chatUsecase.EXPECT().DeleteChatById(context.TODO(), uint64(1)).Return(nil).Times(1)

	err := handler.DeleteChatHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}
