package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"project/internal/auth/usecase/mocks"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"testing"
)

type testCase struct {
	body   []byte
	status int
	name   string
}

func TestHandlers_Signup_Created(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusCreated,
		"Successful registration"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, nil).Times(1)
	usecase.EXPECT().CreateSessionById(ctx, user.Id).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_EmailRegistered(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusConflict,
		"This email is already in the database"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, []error{myErrors.ErrEmailIsAlreadyRegistred}).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_UsernameRegistered(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusConflict,
		"This username is already in the database"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, []error{myErrors.ErrUsernameIsAlreadyRegistred}).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_InvalidEmail(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid email"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, []error{myErrors.ErrInvalidEmail}).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_InvalidUsername(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid username"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, []error{myErrors.ErrInvalidUsername}).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_InvalidPassword(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid password"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Signup(ctx, user).Return(user, []error{myErrors.ErrInvalidPassword}).Times(1)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Signup_InternalError(t *testing.T) {
	test := testCase{[]byte(``),
		http.StatusInternalServerError,
		"Empty body json error"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)

	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.SignupHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Login_OK(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusOK,
		"Successfull login"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Login(ctx, user).Return(user, nil).Times(1)
	usecase.EXPECT().CreateSessionById(ctx, user.Id).Return(model.Session{}, nil).Times(1)

	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Login_UserNotFound(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusNotFound,
		"User not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Login(ctx, user).Return(user, myErrors.ErrUserNotFound).Times(1)

	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Login_IncorrectPassword(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusNotFound,
		"Incorrect password"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	usecase.EXPECT().Login(ctx, user).Return(user, myErrors.ErrIncorrectPassword).Times(1)

	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Login_InternalError(t *testing.T) {
	test := testCase{[]byte(``),
		http.StatusInternalServerError,
		"Empty body json error"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)

	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Auth_UserOK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusOK,
		"User is authorizated"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	cookie := uuid.New().String()

	usecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	usecase.EXPECT().GetUserById(ctx, 1).Return(model.User{}, nil).Times(1)

	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})
	w := httptest.NewRecorder()

	handler.AuthHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Auth_SessionOK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session is exist, user not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	cookie := uuid.New().String()

	usecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	usecase.EXPECT().GetUserById(ctx, 1).Return(model.User{}, myErrors.ErrUserNotFound).Times(1)

	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})
	w := httptest.NewRecorder()

	handler.AuthHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Auth_CookieNotExist(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusUnauthorized,
		"Cookie not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)

	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.AuthHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Auth_SessionNotFound(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	cookie := uuid.New().String()

	usecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{}, myErrors.ErrSessionNotFound).Times(1)

	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})
	w := httptest.NewRecorder()

	handler.AuthHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Logout_Delete(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNoContent,
		"User successfull logout"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	cookie := uuid.New().String()

	usecase.EXPECT().DeleteSessionByCookie(ctx, cookie).Times(1)

	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})
	w := httptest.NewRecorder()

	handler.LogoutHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Logout_CookieNotExist(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusUnauthorized,
		"Cookie not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)

	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	w := httptest.NewRecorder()

	handler.LogoutHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}

func TestHandlers_Logout_SessionNotFound(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	router := mux.NewRouter()
	usecase := mocks.NewMockUsecase(ctl)
	handler := NewAuthHandler(router, usecase)
	ctx := context.Background()
	cookie := uuid.New().String()

	usecase.EXPECT().DeleteSessionByCookie(ctx, cookie).Return(myErrors.ErrSessionNotFound).Times(1)

	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})
	w := httptest.NewRecorder()

	handler.LogoutHandler(w, r)

	require.Equal(t, test.status, w.Code, test.name)
}
