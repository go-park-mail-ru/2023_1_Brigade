package http_utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const cookieTTL = 10 * time.Hour

type jsonErrors struct {
	Err error
}

type jsonUserId struct {
	userID uint64 `json:"user_id"`
}

func (j jsonErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Err.Error())
}

func setHeader(w http.ResponseWriter, err error) {
	for _, item := range error2HttpCode {
		if errors.Is(err, item.Error) {
			w.WriteHeader(item.HttpCode)
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func writeInWriter(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)

	if err != nil {
		log.Error(err)
	}
}

func JsonWriteUserId(w http.ResponseWriter, userID uint64) {
	jsonId, err := json.Marshal(jsonUserId{userID: userID})

	if err != nil {
		setHeader(w, err)
		log.Error(err)
		return
	}

	setHeader(w, myErrors.UserIdGiven)
	writeInWriter(w, jsonId)
}

func JsonWriteUser(w http.ResponseWriter, user model.User) {
	jsonUser, err := json.Marshal(user)

	if err != nil {
		setHeader(w, err)
		log.Error(err)
		return
	}

	setHeader(w, myErrors.UserCreated)
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
		Name:    "session_id",
		Value:   session.Cookie,
		Expires: time.Now().Add(cookieTTL),
	}
	http.SetCookie(w, cookie)
}
