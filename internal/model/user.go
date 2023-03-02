package model

type User struct {
	Id       uint64 `json:"id"       valid:"type(int)"`
	Username string `json:"username" valid:"type(string), required"`
	Name     string `json:"name"     valid:"type(string), required"`
	Email    string `json:"email"    valid:"email, required"`
	Status   string `json:"status"   valid:"type(string)"`
	Password string `json:"password" valid:"type(string),required"`
}

//func a() {
//	govalidator
//}
//govalidator.

//govalidator.CustomTypeTagMap.Set("customMinLengthValidator", func(i interface{}, context interface{}) bool {
//	switch v := context.(type) { // this validates a field against the value in another field, i.e. dependent validation
//		case StructWithCustomByteArray:
//		return len(v.ID) >= v.CustomMinLength
//	}
//	return false
//})
