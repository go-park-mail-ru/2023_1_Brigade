package security

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"project/internal/model"
)

func Hash(password string) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(password))

	if err != nil {
		return "", err
	}

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha, nil
}

func setUserValidators() {
	govalidator.CustomTypeTagMap.Set("usernameValidator", func(i interface{}, context interface{}) bool {
		return i.(string) != ""
	})

	govalidator.CustomTypeTagMap.Set("emailValidator", func(i interface{}, context interface{}) bool {
		return govalidator.IsEmail(i.(string))
	})

	govalidator.CustomTypeTagMap.Set("passwordValidator", func(i interface{}, context interface{}) bool {
		return i.(string) != ""
	})
}

func ValidateSignup(user model.User) []error {
	setUserValidators()

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		errors := err.(govalidator.Errors).Errors()
		return errors
	}

	return nil
}
