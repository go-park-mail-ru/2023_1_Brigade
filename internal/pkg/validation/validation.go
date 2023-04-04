package validation

import (
	"github.com/asaskevich/govalidator"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"strings"
)

func ErrorConversion(err error) error {
	words := strings.Split(err.Error(), " ")
	switch words[0] {
	case "username:":
		return myErrors.ErrInvalidUsername
	case "nickname:":
		return myErrors.ErrInvalidNickname
	case "email:":
		return myErrors.ErrInvalidEmail
	case "password:":
		return myErrors.ErrInvalidPassword
	}

	return nil
}

func setUserValidators() {
	govalidator.CustomTypeTagMap.Set("usernameValidator", func(i interface{}, context interface{}) bool {
		return len(i.(string)) > 7
	})

	govalidator.CustomTypeTagMap.Set("nicknameValidator", func(i interface{}, context interface{}) bool {
		return len(i.(string)) > 7
	})

	govalidator.CustomTypeTagMap.Set("emailValidator", func(i interface{}, context interface{}) bool {
		return govalidator.IsEmail(i.(string))
	})

	govalidator.CustomTypeTagMap.Set("passwordValidator", func(i interface{}, context interface{}) bool {
		return len(i.(string)) > 7
	})
}

func ValidateUser(user model.AuthorizedUser) []error {
	setUserValidators()

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		errors := err.(govalidator.Errors).Errors()
		return errors
	}

	return nil
}
