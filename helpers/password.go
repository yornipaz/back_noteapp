package helpers

import "golang.org/x/crypto/bcrypt"

/*
	CreatePassword

Function  encrypt password using bcrypt library and returns password encrypted

	@params  password string
	@return hashPassword []byte, err error
*/
func CreatePassword(password string) (hashPassword []byte, err error) {
	hashPassword, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	return
}

/*
	ValidatePassword

Function  validated password using password library bcrypt

	@params  password string, userPassword string
	@return err error
*/
func ValidatePassword(password string, userPassword string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return
}
