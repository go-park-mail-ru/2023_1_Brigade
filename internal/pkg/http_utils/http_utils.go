package http_utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Response struct {
	Status int    `json:"status"`
	Data   []byte `json:"data"`
}

const (
	STATUS_OK           = 0
	STATUS_CREATED      = 1
	STATUS_DELETED      = 2
	STATUS_REDIRECTED   = 3
	STATUS_UNAUTHORIZED = 4
	STATUS_NOT_FOUND    = 5
	STATUS_INTERNAL_ERR = 6
)

func SendJsonResponse(w http.ResponseWriter, response Response) {
	switch response.Status {
	case STATUS_OK:
		w.WriteHeader(http.StatusOK)
	case STATUS_CREATED:
		w.WriteHeader(http.StatusCreated)
	case STATUS_DELETED:
		w.WriteHeader(http.StatusNoContent)
	case STATUS_REDIRECTED:
		w.WriteHeader(http.StatusMovedPermanently)
	case STATUS_UNAUTHORIZED:
		w.WriteHeader(http.StatusUnauthorized)
	case STATUS_NOT_FOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response.Data)
}

func ParsingIdUrl(r *http.Request, param string) int {
	vars := mux.Vars(r)
	entitiesID, _ := strconv.Atoi(vars[param])

	return entitiesID
}
