package model

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"         db:"id"`
	Username string `json:"username" valid:"usernameValidator" db:"username"`
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Status   string `json:"status"   valid:"type(string)"      db:"status"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type LoginUser struct {
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type RegistrationUser struct {
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Email    string `json:"email"    valid:"emailValidator"    db:"email"`
	Password string `json:"password" valid:"passwordValidator" db:"password"`
}

type UpdateUser struct {
	Username        string `json:"username"         valid:"usernameValidator" db:"username"`
	Email           string `json:"email"            valid:"emailValidator"    db:"email"`
	Status          string `json:"status"           valid:"type(string)"      db:"status"`
	CurrentPassword string `json:"current_password" valid:"passwordValidator" db:"current_password"`
	NewPassword     string `json:"new_password"     valid:"passwordValidator" db:"new_password"`
}

type Contact struct {
	Username string `json:"username" valid:"usernameValidator" db:"username"`
	Nickname string `json:"nickname" valid:"nicknameValidator" db:"nickname"`
	Status   string `json:"status"   valid:"type(string)"      db:"status"`
	Avatar   string `json:"avatar"                             db:"avatar"`
}

type UserContact struct {
	IdUser    uint64 `json:"id_user"    db:"id_user"`
	IdContact uint64 `json:"id_contact" db:"id_contact"`
}

type ImageUrl struct {
	IdImage  uint64 `json:"id_image"  db:"id_image"`
	ImageUrl string `json:"image_url" db:"image_url"`
}

type UserAvatar struct {
	IdUser  uint64 `json:"id_user"  db:"id_user"`
	IdImage uint64 `json:"id_image" db:"id_image"`
}
