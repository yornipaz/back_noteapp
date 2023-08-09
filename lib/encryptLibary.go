package lib

import "golang.org/x/crypto/bcrypt"

type IEncryptLibrary interface {
	Create(password string, cost int) (hash []byte, err error)
	Validate(password string, hashPassword string) (err error)
}

type EncryptLibrary struct {
}

/*
	Create

Method  encrypt password library and returns password encrypted

	@params  password string
	@params  cost int
	@return hash []byte, err error
*/
func (*EncryptLibrary) Create(password string, cost int) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), cost)
	return
}

/*
	Validate

Method  validated password

	@params  password string, hastPassword string
	@return err error
*/
func (*EncryptLibrary) Validate(password string, hashPassword string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return
}

func NewEncryptLibrary() IEncryptLibrary {
	return &EncryptLibrary{}
}
