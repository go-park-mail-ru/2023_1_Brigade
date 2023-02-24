package model

type User struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
