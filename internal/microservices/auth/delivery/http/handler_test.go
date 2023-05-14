package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	authUserMock "project/internal/microservices/auth/usecase/mocks"
	userMock "project/internal/microservices/user/usecase/mocks"
	"project/internal/model"
	authSessionMock "project/internal/monolithic_services/session/usecase/mocks"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/http_utils"
	"testing"
)

type testCase struct {
	body   []byte
	status int
	name   string
}

func TestHandlers_Signup_Created(t *testing.T) {
	test := testCase{[]byte(`{"email":"marcussss1@mail.ru",
								   "nickname":"marcussss1", 
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, nil).Times(1)
	authSessionUsecase.EXPECT().CreateSessionById(context.TODO(), user.Id).Times(1)

	err = handler.SignupHandler(ctx)
	require.NoError(t, err)
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

	e := echo.New()
	r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, myErrors.ErrEmailIsAlreadyRegistered).Times(1)

	err = handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, myErrors.ErrUsernameIsAlreadyRegistered).Times(1)

	err = handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, myErrors.ErrInvalidEmail).Times(1)

	err = handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, myErrors.ErrInvalidUsername).Times(1)

	err = handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	var registrationUser model.RegistrationUser
	var user model.User

	err := json.Unmarshal(test.body, &registrationUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Signup(context.TODO(), registrationUser).Return(user, myErrors.ErrInvalidPassword).Times(1)

	err = handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)

	err := handler.SignupHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	loginUser := model.LoginUser{}
	user := model.User{}

	err := json.Unmarshal(test.body, &loginUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Login(context.TODO(), loginUser).Return(user, nil).Times(1)
	authSessionUsecase.EXPECT().CreateSessionById(context.TODO(), user.Id).Return(model.Session{}, nil).Times(1)

	err = handler.LoginHandler(ctx)
	require.NoError(t, err)
	require.Equal(t, test.status, w.Code, test.name)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	loginUser := model.LoginUser{}
	user := model.User{}

	err := json.Unmarshal(test.body, &loginUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Login(context.TODO(), loginUser).Return(user, myErrors.ErrEmailNotFound).Times(1)

	err = handler.LoginHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	loginUser := model.LoginUser{}
	user := model.User{}

	err := json.Unmarshal(test.body, &loginUser)
	require.NoError(t, err)

	authUserUsecase.EXPECT().Login(context.TODO(), loginUser).Return(user, myErrors.ErrIncorrectPassword).Times(1)

	err = handler.LoginHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)

	err := handler.LoginHandler(ctx)
	require.Equal(t, http_utils.StatusCode(err), test.status)
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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authSessionUsecase.EXPECT().GetSessionByCookie(context.TODO(), cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	userUsecase.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(model.User{}, nil).Times(1)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authSessionUsecase.EXPECT().GetSessionByCookie(context.TODO(), cookie).Return(model.Session{UserId: 1, Cookie: cookie}, nil).Times(1)
	userUsecase.EXPECT().GetUserById(context.TODO(), uint64(1)).Return(model.User{}, myErrors.ErrUserNotFound).Times(1)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authSessionUsecase.EXPECT().GetSessionByCookie(context.TODO(), cookie).Return(model.Session{}, myErrors.ErrSessionNotFound).Times(1)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authSessionUsecase.EXPECT().DeleteSessionByCookie(context.TODO(), cookie).Times(1)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)

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

	authUserUsecase := authUserMock.NewMockUsecase(ctl)
	authSessionUsecase := authSessionMock.NewMockUsecase(ctl)
	userUsecase := userMock.NewMockUsecase(ctl)
	handler := NewAuthHandler(e, authUserUsecase, authSessionUsecase, userUsecase)
	r.AddCookie(&http.Cookie{Name: "session_id", Value: cookie})

	authSessionUsecase.EXPECT().DeleteSessionByCookie(context.TODO(), cookie).Return(myErrors.ErrSessionNotFound).Times(1)

	err := handler.LogoutHandler(ctx)

	require.Equal(t, http_utils.StatusCode(err), test.status)
}
