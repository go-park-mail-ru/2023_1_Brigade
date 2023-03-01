package security

import (
	"crypto/sha1"
	"encoding/base64"
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
