package security

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

var salt = []byte{0x5C, 0x72, 0x69, 0x67, 0x61, 0x64, 0x65}

func Hash(password []byte) string {
	hashedPassword := argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
	hashString := base64.StdEncoding.EncodeToString(hashedPassword)
	return hashString
}
