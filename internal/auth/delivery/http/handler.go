package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string `json:"message"`
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body errorResponse true "Create order"
// @Success 200 {object} errorResponse
// @Router /orders [post]
func Auth(w http.ResponseWriter, r *http.Request) {
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body errorResponse true "Create order"
// @Success 200 {object} errorResponse
// @Router /orders [post]
func Login(w http.ResponseWriter, r *http.Request) {
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body errorResponse true "Create order"
// @Success 200 {object} errorResponse
// @Router /orders [post]
func Logout(w http.ResponseWriter, r *http.Request) {
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body errorResponse true "Create order"
// @Success 200 {object} errorResponse
// @Router /orders [post]
func Signup(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler(r *mux.Router) {
	r.HandleFunc("/auth", Auth).Methods("PUT")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")
	r.HandleFunc("/signup", Signup).Methods("POST")
}
