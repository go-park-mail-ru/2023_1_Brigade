package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	authMock "project/internal/auth/usecase/mocks"
	myMiddleware "project/internal/middleware"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/http_utils"
	userMock "project/internal/user/usecase/mocks"
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

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, nil).Times(1)
	authUsecase.EXPECT().CreateSessionById(ctx, user.Id).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err = handler.SignupHandler(ctx)
		require.NoError(t, err)
		require.Equal(t, test.status, w.Code, test.name)
		return err
	})(ctx)
}

func TestHandlers_Signup_EmailRegistered(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusConflict,
		"This email is already in the database"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, myErrors.ErrEmailIsAlreadyRegistred).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Signup_UsernameRegistered(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusConflict,
		"This username is already in the database"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, myErrors.ErrUsernameIsAlreadyRegistred).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Signup_InvalidEmail(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid email"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, myErrors.ErrInvalidEmail).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Signup_InvalidUsername(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid username"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, myErrors.ErrInvalidUsername).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Signup_InvalidPassword(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
						  "username":"marcussss1",
						  "password":"baumanka"}`),
		http.StatusBadRequest,
		"Invalid password"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Signup(ctx, user).Return(user, myErrors.ErrInvalidPassword).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Signup_InternalError(t *testing.T) {
	test := testCase{[]byte(`{sfadfad{f`),
		http.StatusInternalServerError,
		"Empty body json error"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)

	_ = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.SignupHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Login_OK(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusOK,
		"Successfull login"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Login(ctx, user).Return(user, nil).Times(1)
	authUsecase.EXPECT().CreateSessionById(ctx, user.Id).Return(model.Session{}, nil).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err = handler.LoginHandler(ctx)
		require.NoError(t, err)
		require.Equal(t, test.status, w.Code, test.name)
		return err
	})(ctx)
}

func TestHandlers_Login_UserNotFound(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusNotFound,
		"User not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Login(ctx, user).Return(user, myErrors.ErrUserNotFound).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.LoginHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Login_IncorrectPassword(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "password":"baumanka"}`),
		http.StatusNotFound,
		"Incorrect password"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	user := model.User{}

	err := json.Unmarshal(test.body, &user)
	require.NoError(t, err)

	authUsecase.EXPECT().Login(ctx, user).Return(user, myErrors.ErrIncorrectPassword).Times(1)

	err = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.LoginHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Login_InternalError(t *testing.T) {
	test := testCase{[]byte(`{adgadgadg{`),
		http.StatusInternalServerError,
		"Empty body json error"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)

	_ = myMiddleware.XSSMidlleware(func(ctx echo.Context) error {
		err := handler.LoginHandler(ctx)
		require.Equal(t, http_utils.StatusCode(err), test.status)
		return err
	})(ctx)
}

func TestHandlers_Auth_UserOK(t *testing.T) {
	test := testCase{[]byte(``),
		http.StatusOK,
		"User is authorizated"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	cookie := uuid.New().String()

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authUsecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	userUsecase.EXPECT().GetUserById(ctx, 1).Return(model.User{}, nil).Times(1)

	err := handler.AuthHandler(ctx)

	require.NoError(t, err)
}

func TestHandlers_Auth_SessionOK(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session is exist, user not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	cookie := uuid.New().String()

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authUsecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	userUsecase.EXPECT().GetUserById(ctx, 1).Return(model.User{}, myErrors.ErrUserNotFound).Times(1)

	err := handler.AuthHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}

func TestHandlers_Auth_CookieNotExist(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusUnauthorized,
		"Cookie not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)

	err := handler.AuthHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}

func TestHandlers_Auth_SessionNotFound(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	cookie := uuid.New().String()

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authUsecase.EXPECT().GetSessionByCookie(ctx, cookie).Return(model.Session{}, myErrors.ErrSessionNotFound).Times(1)

	err := handler.AuthHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}

func TestHandlers_Logout_Delete(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNoContent,
		"User successfull logout"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	cookie := uuid.New().String()

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authUsecase.EXPECT().DeleteSessionByCookie(ctx, cookie).Times(1)

	err := handler.LogoutHandler(ctx)

	require.NoError(t, err)
}

func TestHandlers_Logout_CookieNotExist(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusUnauthorized,
		"Cookie not exist"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)

	err := handler.LogoutHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}

func TestHandlers_Logout_SessionNotFound(t *testing.T) {
	test := testCase{[]byte(""),
		http.StatusNotFound,
		"Session not found"}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	e := echo.New()
	r := httptest.NewRequest("DELETE", "/logout/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	cookie := uuid.New().String()

	authUsecase := authMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authUsecase.EXPECT().DeleteSessionByCookie(ctx, cookie).Return(myErrors.ErrSessionNotFound).Times(1)

	err := handler.LogoutHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}
