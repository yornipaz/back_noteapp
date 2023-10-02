package authrepository

import (
	"os"

	"github.com/yornifpaz/back_noteapp/lib"
)

type IAuthRepository interface {
	CreateJWT(ID string, Payload map[string]interface{}) (tokenString string, err error)
	ValidatePassword(password string, userPassword string) (err error)
	CreatePassword(password string) (hashPassword []byte, err error)
}
type AuthRepository struct {
	jwtLibrary     lib.IJwtLibrary
	encryptLibrary lib.IEncryptLibrary
}

/*
	CreateJWT

Method create a new token object, specifying signing method and the claims

	@param  ID unit identifier  user
	@return tokenString string and error error
*/
func (repository *AuthRepository) CreateJWT(ID string, Payload map[string]interface{}) (tokenString string, err error) {
	claims := lib.JwtClaimsCustom{
		Id:      ID,
		Expired: 24,
		Secret:  os.Getenv("SECRET_KEY"),
		Payload: Payload,
	}
	tokenString, err = repository.jwtLibrary.Create(claims)
	return
}

// CreatePassword implements IAuthRepository.
func (repository *AuthRepository) CreatePassword(password string) (hashPassword []byte, err error) {
	hashPassword, err = repository.encryptLibrary.Create(password, 10)
	return
}

// ValidatePassword implements IAuthRepository.
func (repository *AuthRepository) ValidatePassword(password string, userPassword string) (err error) {
	err = repository.encryptLibrary.Validate(password, userPassword)
	return
}

func NewAuthRepository(jwtLibrary lib.IJwtLibrary, encryptLibrary lib.IEncryptLibrary) IAuthRepository {
	return &AuthRepository{jwtLibrary: jwtLibrary, encryptLibrary: encryptLibrary}
}
