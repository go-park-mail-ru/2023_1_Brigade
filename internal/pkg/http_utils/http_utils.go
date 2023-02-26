package http_utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ParsingIdUrl(r *http.Request, param string) (int, error) {
	vars := mux.Vars(r)
	entitiesID, err := strconv.Atoi(vars[param])

	return entitiesID, err
}
