package model

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Password string `json:"password"`
}
