package security

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"project/internal/model"
	"project/internal/pkg/errors"
)

//const minNicknameLength = 4
//const maxNicknameLength = 20
//
//const minPasswordLength = 6
//const maxPasswordLength = 20
//
//func ValidateNickname(nickname string) error {
//	len := len(nickname)
//	if len < minNicknameLength || len > maxNicknameLength {
//		return errors.ErrInvalidNickname
//	}
//	return nil
//}
//
//func ValidateEmail(email string) error {
//	if _, err := mail.ParseAddress(email); err != nil {
//		return errors.ErrInvalidEmail
//	}
//	return nil
//}
//
//func ValidatePassword(password string) error {
//	len := len(password)
//	if len < minPasswordLength || len > maxPasswordLength {
//		return errors.ErrInvalidPassword
//	}
//	return nil
//}

func Hash(password string) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(password))

	if err != nil {
		return "", err
	}

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha, nil
}

type validateUsername struct {
	Username string `valid:"usernameValidator"`
}

type validateName struct {
	Name string `valid:"nameValidator"`
}

type validateEmail struct {
	Email string `valid:"emailValidator"`
}

type validatePassword struct {
	Password string `valid:"passwordValidator"`
}

func setUsernameValidator() {
	govalidator.CustomTypeTagMap.Set("usernameValidator", func(i interface{}, context interface{}) bool {
		return i.(string) != ""
	})
}

func setNameValidator() {
	govalidator.CustomTypeTagMap.Set("nameValidator", func(i interface{}, context interface{}) bool {
		return i.(string) != ""
	})
}

func setEmailValidator() {
	govalidator.CustomTypeTagMap.Set("emailValidator", func(i interface{}, context interface{}) bool {
		return govalidator.IsEmail(i.(string))
	})
}

func setPasswordValidator() {
	govalidator.CustomTypeTagMap.Set("passwordValidator", func(i interface{}, context interface{}) bool {
		return i.(string) != ""
	})
}

//type jsonErrors struct {
//	error
//}
//
//func (e jsonErrors) MarshalJSON() ([]byte, error) {
//	return json.Marshal(e.Error())
//}

func ValidateSignup(user model.User) []error {
	var validateErrors []error

	setUsernameValidator()
	setNameValidator()
	setEmailValidator()
	setPasswordValidator()

	username := validateUsername{user.Username}
	name := validateName{user.Name}
	email := validateEmail{user.Email}
	password := validatePassword{user.Password}

	usernameIsValidate, _ := govalidator.ValidateStruct(username)
	nameIsValidate, _ := govalidator.ValidateStruct(name)
	emailIsValidate, _ := govalidator.ValidateStruct(email)
	passwordIsValidate, _ := govalidator.ValidateStruct(password)

	//fmt.Println(usernameIsValidate, nameIsValidate, emailIsValidate, passwordIsValidate)
	//switch {
	//case !usernameIsValidate:
	//	validateErrors = append(validateErrors, errors.InvalidUsername)
	//	fallthrough
	//case !nameIsValidate:
	//	validateErrors = append(validateErrors, errors.InvalidName)
	//	fallthrough
	//case !emailIsValidate:
	//	validateErrors = append(validateErrors, errors.InvalidEmail)
	//	fallthrough
	//case !passwordIsValidate:
	//	validateErrors = append(validateErrors, errors.InvalidPassword)
	//}

	// почему-то через кейс какая-то беда, и он password тоже добавляет

	if !usernameIsValidate {
		validateErrors = append(validateErrors, errors.InvalidUsername)
	}

	if !nameIsValidate {
		validateErrors = append(validateErrors, errors.InvalidName)
	}

	if !emailIsValidate {
		validateErrors = append(validateErrors, errors.InvalidEmail)
	}

	if !passwordIsValidate {
		validateErrors = append(validateErrors, errors.InvalidPassword)
	}

	//if !usernameIsValidate {
	//	validateErrors = append(validateErrors, jsonErrors{errors.InvalidUsername})
	//}
	//
	//if !nameIsValidate {
	//	validateErrors = append(validateErrors, jsonErrors{errors.InvalidName})
	//}
	//
	//if !emailIsValidate {
	//	validateErrors = append(validateErrors, jsonErrors{errors.InvalidEmail})
	//}
	//
	//if passwordIsValidate {
	//	validateErrors = append(validateErrors, jsonErrors{errors.InvalidPassword})
	//}

	//jsonValidateErrors, err := json.Marshal(validateErrors)
	return validateErrors
}
