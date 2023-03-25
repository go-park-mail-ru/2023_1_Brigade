package http

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	myMiddleware "project/internal/middleware"
	"project/internal/model"
	userMock "project/internal/user/usecase/mocks"
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

	userUsecase.EXPECT().GetUserById(ctx, 1).Return(model.User{}, nil).Times(1)

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

	userUsecase.EXPECT().GetUserById(ctx, 1).Return(model.User{Id: 1}, nil).Times(1)

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

	userUsecase.EXPECT().DeleteUserById(ctx, 1).Return(nil).Times(1)

	err := handler.DeleteUserHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}

func TestHandlers_PutUser_OK(t *testing.T) {
	test := testCase{[]byte(`{
		"username":        "marcussss",
		"email":           "marcussss@gmail.com",
		"status":          "I'm marcussss",
		"current_password": "baumanka",
		"new_password":     "baumanka_cool"
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
		Username:        "marcussss",
		Email:           "marcussss@gmail.com",
		Status:          "I'm marcussss",
		CurrentPassword: "baumanka",
		NewPassword:     "baumanka_cool",
	}
	user := model.User{
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

	userUsecase.EXPECT().PutUserById(ctx, newUser, 1).Return(user, nil).Times(1)

	_ = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.PutUserHandler(ctx)

		require.NoError(t, err)
		require.Equal(t, test.status, w.Code)

		return err
	})(ctx)
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

	userUsecase.EXPECT().GetUserContacts(ctx, 1).Return([]model.User{}, nil).Times(1)

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

	userUsecase.EXPECT().AddUserContact(ctx, 1, 2).Return(model.User{Id: 2}, nil).Times(1)

	err := handler.UserAddContactHandler(ctx)

	require.NoError(t, err)
	require.Equal(t, test.status, w.Code)
}
