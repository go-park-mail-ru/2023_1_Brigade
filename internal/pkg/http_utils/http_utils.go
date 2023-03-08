package http_utils

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strconv"
	"strings"
	"time"
)

type jsonErrors struct {
	Err error
}

func (j jsonErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Err.Error())
}

func ErrorsConversion(validateErrors []error) []error {
	var errors []error
	for _, err := range validateErrors {
		words := strings.Split(err.Error(), " ")
		switch words[0] {
		case "username:":
			errors = append(errors, myErrors.ErrInvalidUsername)
		case "email:":
			errors = append(errors, myErrors.ErrInvalidEmail)
		case "password:":
			errors = append(errors, myErrors.ErrInvalidPassword)
		}
	}
	return errors
}

func setHeader(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, myErrors.UserGetting):
		w.WriteHeader(http.StatusOK)
	case errors.Is(err, myErrors.UserCreated):
		w.WriteHeader(http.StatusCreated)
	case errors.Is(err, myErrors.SessionSuccessDeleted):
		w.WriteHeader(http.StatusNoContent)
	case errors.Is(err, myErrors.ErrInvalidUsername):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, myErrors.ErrInvalidEmail):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, myErrors.ErrInvalidPassword):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, myErrors.ErrEmailIsAlreadyRegistred):
		w.WriteHeader(http.StatusConflict)
	case errors.Is(err, myErrors.ErrUsernameIsAlreadyRegistred):
		w.WriteHeader(http.StatusConflict)
	case errors.Is(err, myErrors.ErrSessionIsAlreadyCreated):
		w.WriteHeader(http.StatusConflict)
	case errors.Is(err, myErrors.ErrCookieNotFound):
		w.WriteHeader(http.StatusUnauthorized)
	case errors.Is(err, myErrors.ErrSessionNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, myErrors.ErrUserNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, myErrors.ErrIncorrectPassword):
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeInWriter(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)

	if err != nil {
		log.Error(err)
	}
}

func JsonWriteUserCreated(w http.ResponseWriter, user model.User) {
	jsonUser, err := json.Marshal(user)

	if err != nil {
		setHeader(w, err)
		log.Error(err)
		return
	}

	setHeader(w, myErrors.UserCreated)
	writeInWriter(w, jsonUser)
}

func JsonWriteUserGet(w http.ResponseWriter, user model.User) {
	jsonUser, err := json.Marshal(user)

	if err != nil {
		setHeader(w, err)
		log.Error(err)
		return
	}

	setHeader(w, myErrors.UserGetting)
	writeInWriter(w, jsonUser)
}

func JsonWriteErrors(w http.ResponseWriter, errors []error) {

	for _, err := range errors {
		log.Error(err)
	}

	var JsonErrors []jsonErrors
	for _, err := range errors {
		JsonErrors = append(JsonErrors, jsonErrors{Err: err}) // если ошибка валидации выдаст сразу несколько
	}

	jsonValidateErrors, err := json.Marshal(JsonErrors)

	if err != nil {
		setHeader(w, err)
		log.Error(err)
		return
	}

	setHeader(w, errors[0])
	writeInWriter(w, jsonValidateErrors)
}

func ParsingIdUrl(r *http.Request, param string) (uint64, error) {
	vars := mux.Vars(r)
	entitiesID, err := strconv.ParseUint(vars[param], 10, 64)

	if err != nil {
		return 0, err
	}

	return entitiesID, nil
}

func SetCookie(w http.ResponseWriter, session model.Session) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    session.Cookie,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,                  // local
		SameSite: http.SameSiteNoneMode, // local
		Expires:  time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		HttpOnly: true,
		Secure:   true,                  // local
		SameSite: http.SameSiteNoneMode, // local
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
	})
}
