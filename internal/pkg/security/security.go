package security

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"project/internal/model"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GenerateFilename(userID uint64, filename string) string {
	filename = reverse(filename)
	extensionFilename := strings.Split(filename, ".")
	extension := reverse(extensionFilename[0])
	generatedFilename := strconv.FormatUint(userID, 10) + "." + extension
	return generatedFilename
}

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
		return len(i.(string)) > 1
	})

	govalidator.CustomTypeTagMap.Set("emailValidator", func(i interface{}, context interface{}) bool {
		return govalidator.IsEmail(i.(string))
	})

	govalidator.CustomTypeTagMap.Set("passwordValidator", func(i interface{}, context interface{}) bool {
		return len(i.(string)) > 8
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
