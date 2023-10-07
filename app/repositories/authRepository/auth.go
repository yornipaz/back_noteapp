package authrepository

import (
	"fmt"

	"github.com/yornifpaz/back_noteapp/lib"
)

type IAuthRepository interface {
	CreateJWT(ID string, expired int64, secret string, Payload map[string]interface{}) (tokenString string, err error)
	ValidatePassword(password string, userPassword string) (err error)
	CreatePassword(password string) (hashPassword []byte, err error)
	ValidateToken(token string) (id string, err error)
}
type AuthRepository struct {
	jwtLibrary     lib.IJwtLibrary
	encryptLibrary lib.IEncryptLibrary
}

// ValidateToken implements IAuthRepository.
func (repository *AuthRepository) ValidateToken(token string) (id string, err error) {
	isValidToken := repository.jwtLibrary.ValidateToken(token)
	if !isValidToken {
		return "", fmt.Errorf("el token no es válido o ya expiró")
	}
	tokenString, err := repository.jwtLibrary.Parse(token)
	id = repository.jwtLibrary.GetClaims(tokenString)["sub"].(string)
	return
}

/*
	CreateJWT

Method create a new token object, specifying signing method and the claims

	@param  ID unit identifier  user
	@return tokenString string and error error
*/
func (repository *AuthRepository) CreateJWT(ID string, expired int64, secret string, Payload map[string]interface{}) (tokenString string, err error) {
	claims := lib.JwtClaimsCustom{
		Id:      ID,
		Expired: expired,
		Secret:  secret,
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
