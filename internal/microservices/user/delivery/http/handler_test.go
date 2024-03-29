package http

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	userMock "project/internal/microservices/user/usecase/mocks"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	"testing"
)

type testCase struct {
	body   []byte
	status int
	name   string
}

func TestHandlers_GetUser_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusOK,
		"Successfull getting user"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/1", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.SetParamNames("userID")
	ctx.SetParamValues("1")

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(model.User{}, nil).Times(1)

	err := handler.GetUserHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_GetCurrentUser_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusOK,
		"Successfull get current user"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/settings", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(model.User{Id: 1}, nil).Times(1)

	err := handler.GetCurrentUserHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_DeleteUser_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNoContent,
		"Successfull delete user"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/remove", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})
	ctx.SetParamNames("userID")
	ctx.SetParamValues("1")

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().DeleteUserById(context.TODO(), uint64(1)).Return(nil).Times(1)

	err := handler.DeleteUserHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_PutUserInfo_OK(t *testing.T) {
	test := testCase{[]byte(`{
		"nickname":        "marcussss",
		"email":           "marcussss@gmail.com",
		"status":          "I'm marcussss"
	}`),
		http.StatusOK,
		"Successfull put user"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/settings", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	newUser := model.UpdateUser{
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "I'm marcussss",
	}
	user := model.AuthorizedUser{
		Id:       1,
		Username: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "I'm marcussss",
		Password: "baumanka_cool",
	}
	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})
	ctx.SetParamNames("userID")
	ctx.SetParamValues("1")

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().PutUserById(context.TODO(), newUser, uint64(1)).Return(model_conversion.FromAuthorizedUserToUser(user), nil).Times(1)

	err := handler.PutUserHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_GetUserContacts_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusOK,
		"Successfull getting user contacts"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/1/contacts", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})
	ctx.SetParamNames("userID")
	ctx.SetParamValues("1")

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().GetAllUsersExceptCurrentUser(context.TODO(), uint64(1)).Return([]model.User{}, nil).Times(1)

	err := handler.GetUserContactsHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_UserAddContact_OK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusCreated,
		"Successfull adding user in contact"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/api/v1/users/2/add", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	ctx := e.NewContext(r, w)
	ctx.Set("session", model.Session{UserId: 1})
	ctx.SetParamNames("userID")
	ctx.SetParamValues("2")

	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewUserHandler(e, userUsecase)

	userUsecase.EXPECT().AddUserContact(context.TODO(), uint64(1), uint64(2)).Return([]model.User{}, nil).Times(1)

	err := handler.UserAddContactHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}
