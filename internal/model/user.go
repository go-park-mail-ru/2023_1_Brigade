package model

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)" db:"id"`
	Username string `json:"username" valid:"usernameValidator" db:"username"`
	Email    string `json:"email"    valid:"emailValidator" db:"email"`
	Status   string `json:"status"   valid:"type(string)" db:"status"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}
