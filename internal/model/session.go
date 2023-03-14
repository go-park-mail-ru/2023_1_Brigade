package model

type Session struct {
	UserId uint64 `json:"user_id" db:"user_id"`
	Cookie string `json:"cookie" db:"cookie"`
}
