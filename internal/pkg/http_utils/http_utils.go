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
		log.Println(words[0])
		switch words[0] {
		case "username:":
			errors = append(errors, myErrors.ErrInvalidUsername)
		case "name:":
			errors = append(errors, myErrors.ErrInvalidName)
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
	case errors.Is(err, myErrors.ErrInvalidName):
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

	var validateErrors []jsonErrors
	for _, err := range errors {
		validateErrors = append(validateErrors, jsonErrors{Err: err})
	}

	jsonValidateErrors, err := json.Marshal(validateErrors)

	if err != nil {
		setHeader(w, myErrors.ErrInternal)
		log.Error(err)
		return
	}

	setHeader(w, errors[0])
	writeInWriter(w, jsonValidateErrors)
}

func JsonWriteInternalError(w http.ResponseWriter) {
	setHeader(w, myErrors.ErrInternal)

	internalError := jsonErrors{Err: myErrors.ErrInternal}
	jsonInternalError, err := json.Marshal(internalError)

	if err != nil {
		log.Error(err)
		return
	}

	writeInWriter(w, jsonInternalError)
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
		Expires:  time.Now().Add(10000 * time.Hour),
	}
	http.SetCookie(w, cookie)
}
