package pkg

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ParsingIdUrl(r *http.Request, param string) (int, error) {
	vars := mux.Vars(r)
	chatID, err := strconv.Atoi(vars[param])

	return chatID, err
}
