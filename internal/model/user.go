package model

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"`
	Username string `json:"username" valid:"type(string), required"`
	Name     string `json:"name"     valid:"type(string), required"`
	Email    string `json:"email"    valid:"email, required"`
	Status   string `json:"status"   valid:"type(string)"`
	Password string `json:"password" valid:"type(string),required"`
}
