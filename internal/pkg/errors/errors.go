package errors

import "fmt"

var (
	EmailIsAlreadyRegistred = fmt.Errorf("The email is already registered")
	//PasswordHashWriter      = fmt.Errorf("Password hashing error")
)
