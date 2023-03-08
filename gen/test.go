package gen

//
//import (
//	"bytes"
//	"context"
//	"database/sql"
//	"github.com/gorilla/mux"
//	_ "github.com/lib/pq"
//	log "github.com/sirupsen/logrus"
//	"github.com/stretchr/testify/require"
//	"net/http"
//	"net/http/httptest"
//	authrepository "project/internal/auth/repository"
//	authusecase "project/internal/auth/usecase"
//	myErrors "project/internal/pkg/errors"
//	"testing"
//)
//
//var schema = `
//DROP TABLE Profile;
//DROP TABLE Session;
//
//CREATE TABLE Profile (
//    id       serial,
//    username varchar(255),
//    email    varchar(255),
//    status   varchar(255),
//    password varchar(255)
//);
//
//CREATE TABLE Session (
//    user_id integer,
//    cookie  varchar(255)
//);
//`
//
//var schema2 = `
//DROP TABLE Session;
//
//CREATE TABLE Session (
//    user_id integer,
//    cookie  varchar(255)
//);
//`
//
//func TestHandlers_Signup(t *testing.T) {
//	type testCase struct {
//		body   []byte
//		status int
//		name   string
//	}
//
//	testCases := []testCase{
//		{
//			[]byte(`{"email":"marcussss1@mail.ru",
//						  "username":"marcussss1",
//						  "password":"baumanka"}`),
//			http.StatusCreated,
//			"Successful registration",
//		},
//		{
//			[]byte(`{"email":"marcussss1@mail.ru",
//						  "username":"marcussss2",
//						  "password":"baumanka"}`),
//			http.StatusConflict,
//			"This email is already in the database",
//		},
//		{
//			[]byte(`{"email":"marcussss2@mail.ru",
//						  "username":"marcussss1",
//						  "password":"baumanka"}`),
//			http.StatusConflict,
//			"This username is already in the database",
//		},
//		{
//			[]byte(`{"email":"marcussss1",
//						  "username":"marcussss2",
//						  "password":"baumanka"}`),
//			http.StatusBadRequest,
//			"Invalid email",
//		},
//		{
//			[]byte(`{"email":"marcussss1",
//						  "username":"marcussss2",
//						  "password":"bauman"}`),
//			http.StatusBadRequest,
//			"Invalid username",
//		},
//		{
//			[]byte(`{"email":"marcussss1",
//						  "username":"",
//						  "password":"baumanka"}`),
//			http.StatusBadRequest,
//			"Invalid password",
//		},
//	}
//	log.SetFormatter(&log.TextFormatter{
//		FullTimestamp: true,
//	})
//	log.SetReportCaller(true)
//	connStr := "user=brigade password=123 dbname=brigade sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//
//	db.Exec(schema)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r := mux.NewRouter()
//
//	authRepo := authrepository.NewAuthMemoryRepository(db)
//	authUsecase := authusecase.NewAuthUsecase(authRepo)
//	authHandl := NewAuthHandler(r, authUsecase)
//
//	for _, test := range testCases {
//		r := httptest.NewRequest("POST", "/signup/", bytes.NewReader(test.body))
//		w := httptest.NewRecorder()
//
//		authHandl.SignupHandler(w, r)
//
//		require.Equal(t, test.status, w.Code, test.name)
//	}
//}
//
//func TestHandlers_Login(t *testing.T) {
//	type testCase struct {
//		body   []byte
//		status int
//		name   string
//	}
//
//	testCases := []testCase{
//		{
//			[]byte(`{"email":"marcussss1@mail.ru",
//						  "password":"baumanka"}`),
//			http.StatusOK,
//			"Successful login",
//		},
//		{
//			[]byte(`{"email":"marcussss2@mail.ru",
//						  "password":"baumanka"}`),
//			http.StatusNotFound,
//			"Wrong email",
//		},
//		{
//			[]byte(`{"email":"marcussss1@mail.ru",
//						  "password":"baumanka1"}`),
//			http.StatusNotFound,
//			"Wrong password",
//		},
//	}
//	log.SetFormatter(&log.TextFormatter{
//		FullTimestamp: true,
//	})
//	log.SetReportCaller(true)
//	connStr := "user=brigade password=123 dbname=brigade sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//
//	db.Exec(schema2)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r := mux.NewRouter()
//
//	authRepo := authrepository.NewAuthMemoryRepository(db)
//	authUsecase := authusecase.NewAuthUsecase(authRepo)
//	authHandl := NewAuthHandler(r, authUsecase)
//
//	for _, test := range testCases {
//		r := httptest.NewRequest("POST", "/login/", bytes.NewReader(test.body))
//		w := httptest.NewRecorder()
//
//		authHandl.LoginHandler(w, r)
//
//		require.Equal(t, test.status, w.Code, test.name)
//	}
//}
//
//func TestHandlers_Logout(t *testing.T) {
//	type testCase struct {
//		status int
//		name   string
//	}
//
//	notExistingCookieTest := testCase{
//		http.StatusUnauthorized,
//		"Not existing cookie",
//	}
//
//	incorrectCookieTest := testCase{
//		http.StatusNotFound,
//		"Existing but incorrect cookie",
//	}
//
//	successLogoutTest := testCase{
//		http.StatusNoContent,
//		"Successful logout",
//	}
//
//	log.SetFormatter(&log.TextFormatter{
//		FullTimestamp: true,
//	})
//	log.SetReportCaller(true)
//	connStr := "user=brigade password=123 dbname=brigade sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	router := mux.NewRouter()
//
//	authRepo := authrepository.NewAuthMemoryRepository(db)
//	authUsecase := authusecase.NewAuthUsecase(authRepo)
//	authHandl := NewAuthHandler(router, authUsecase)
//
//	r := httptest.NewRequest("POST", "/logout/", bytes.NewReader([]byte("{}")))
//	w := httptest.NewRecorder()
//
//	authHandl.LogoutHandler(w, r)
//	require.Equal(t, notExistingCookieTest.status, w.Code, notExistingCookieTest.name)
//
//	///////////////////////////////////////////////////////////////////////////////////////////
//
//	r = httptest.NewRequest("POST", "/logout/", bytes.NewReader([]byte("{}")))
//	w = httptest.NewRecorder()
//
//	r.AddCookie(&http.Cookie{
//		Name:  "session_id",
//		Value: "abcdefgh",
//	})
//	authHandl.LogoutHandler(w, r)
//	require.Equal(t, incorrectCookieTest.status, w.Code, incorrectCookieTest.name)
//
//	///////////////////////////////////////////////////////////////////////////////////////////
//
//	r = httptest.NewRequest("POST", "/logout/", bytes.NewReader([]byte("{}")))
//	w = httptest.NewRecorder()
//
//	session, err := authRepo.GetSessionById(context.Background(), 1) // т.к. всего 1 пользователь
//	require.Error(t, err, myErrors.ErrSessionIsAlreadyCreated)
//
//	r.AddCookie(&http.Cookie{
//		Name:  "session_id",
//		Value: session.Cookie,
//	})
//	authHandl.LogoutHandler(w, r)
//
//	require.Equal(t, successLogoutTest.status, w.Code, successLogoutTest.name)
//}
//
//func TestHandlers_Auth(t *testing.T) {
//	type testCase struct {
//		status int
//		name   string
//	}
//
//	successAuthTest := testCase{
//		http.StatusOK,
//		"The user is logged in",
//	}
//
//	notExistingCookieTest := testCase{
//		http.StatusUnauthorized,
//		"Not existing cookie",
//	}
//
//	incorrectCookieTest := testCase{
//		http.StatusNotFound,
//		"Existing but incorrect cookie",
//	}
//
//	log.SetFormatter(&log.TextFormatter{
//		FullTimestamp: true,
//	})
//	log.SetReportCaller(true)
//	connStr := "user=brigade password=123 dbname=brigade sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	router := mux.NewRouter()
//
//	authRepo := authrepository.NewAuthMemoryRepository(db)
//	authUsecase := authusecase.NewAuthUsecase(authRepo)
//	authHandl := NewAuthHandler(router, authUsecase)
//
//	r := httptest.NewRequest("GET", "/auth/", bytes.NewReader([]byte("{}")))
//	w := httptest.NewRecorder()
//
//	authHandl.AuthHandler(w, r)
//	require.Equal(t, notExistingCookieTest.status, w.Code, notExistingCookieTest.name)
//
//	/////////////////////////////////////////////////////////////////////////////////////////////
//
//	r = httptest.NewRequest("POST", "/login/", bytes.NewReader([]byte(`{"email":"marcussss1@mail.ru","password":"baumanka"}`)))
//	w = httptest.NewRecorder()
//
//	authHandl.LoginHandler(w, r)
//	require.Equal(t, http.StatusOK, w.Code, "Login")
//
//	/////////////////////////////////////////////////////////////////////////////////////////////
//
//	r = httptest.NewRequest("GET", "/auth/", bytes.NewReader([]byte("{}")))
//	w = httptest.NewRecorder()
//
//	r.AddCookie(&http.Cookie{
//		Name:  "session_id",
//		Value: "abcdefgh",
//	})
//
//	authHandl.AuthHandler(w, r)
//	require.Equal(t, incorrectCookieTest.status, w.Code, incorrectCookieTest.name)
//
//	/////////////////////////////////////////////////////////////////////////////////////////////
//
//	r = httptest.NewRequest("GET", "/auth/", bytes.NewReader([]byte("{}")))
//	w = httptest.NewRecorder()
//
//	session, err := authRepo.GetSessionById(context.Background(), 1) // т.к. всего 1 пользователь
//	require.Error(t, err, myErrors.ErrSessionIsAlreadyCreated)
//
//	r.AddCookie(&http.Cookie{
//		Name:  "session_id",
//		Value: session.Cookie,
//	})
//
//	authHandl.AuthHandler(w, r)
//	require.Equal(t, successAuthTest.status, w.Code, successAuthTest.name)
//}
