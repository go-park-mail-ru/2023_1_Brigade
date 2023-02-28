package pkg

import (
	"net/http"
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
	STATUS_BAD_METHOD   = 6
	STATUS_INTERNAL     = 7
	STATUS_ERR_JSON     = 8
	STATUS_ERR_DB       = 9
)

//var (
//	OK           = Response{Status: STATUS_CREATED}
//	DELETED      = Response{Status: STATUS_DELETED}
//	REDIRECTED   = Response{Status: STATUS_REDIRECTED}
//	UNAUTHORIZED = Response{Status: STATUS_UNAUTHORIZED}
//	NOT_FOUND    = Response{Status: STATUS_NOT_FOUND}
//	BAD_METHOD   = Response{Status: STATUS_BAD_METHOD}
//	INTERNAL     = Response{Status: STATUS_INTERNAL}
//	ERR_DB       = Response{Status: STATUS_ERR_DB}
//	ERR_JSON     = Response{Status: STATUS_ERR_JSON}
//)

func WriteJsonResponse(w http.ResponseWriter, resp Response) {
	switch resp.Status {
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
	case STATUS_BAD_METHOD:
		w.WriteHeader(http.StatusMethodNotAllowed)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//func WriteJsonErr(w http.ResponseWriter, status int, message string) {
//	err := JsonResponse{
//		Status: status,
//	}
//	WriteJsonErrFull(w, &err)
//}
//
//func CreateJsonErr(status int, message string) *JsonResponse {
//	err := JsonResponse{
//		Status: status,
//	}
//	return &err
//}
