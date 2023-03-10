package model

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"`
	Username string `json:"username" valid:"usernameValidator"`
	Email    string `json:"email"    valid:"emailValidator"`
	Status   string `json:"status"   valid:"type(string)"`
	Password string `json:"password" valid:"passwordValidator"`
}
