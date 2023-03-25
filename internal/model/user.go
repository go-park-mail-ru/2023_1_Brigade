package model

type UserContact struct {
	IdUser    uint64 `json:"id_user"    db:"id_user"`
	IdContact uint64 `json:"id_contact" db:"id_contact"`
}

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"         db:"id"`
	Username string `json:"username" valid:"usernameValidator" db:"username"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Status   string `json:"status"   valid:"type(string)"      db:"status"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type UpdateUser struct {
	Username        string `json:"username"         valid:"usernameValidator" db:"username"`
	Email           string `json:"email"            valid:"emailValidator"    db:"email"`
	Status          string `json:"status"           valid:"type(string)"      db:"status"`
	CurrentPassword string `json:"current_password" valid:"passwordValidator" db:"current_password"`
	NewPassword     string `json:"new_password"     valid:"passwordValidator" db:"new_password"`
}
