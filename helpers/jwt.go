package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
*@package helpers
* function jwt created and validates token
 */

/*
	CreateJWT

function create a new token object, specifying signing method and the claims

	@param  ID unit identifier  user
	@return tokenString string and error error
*/
func CreateJWT(ID string) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return
}

var currentToken *jwt.Token

/*
	IsValidToken

function validates if the token is valid or not expired

	@param  tokenString string
	@return isvalidtoken bool
*/
func IsValidToken(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false
	}
	currentToken = token

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return false
		}

		return true

	} else {
		return false
	}
}

/*
	GetClaims

function  get claims from the current token

	@return claims jwt.MapClaims
*/
func GetClaims() (claims jwt.MapClaims) {
	claims = currentToken.Claims.(jwt.MapClaims)
	return
}

/*
	GetCurrentUserId

function  get  current user id  from the current token

	@return userId string
*/
func GetCurrentUserId() (userId string) {
	userId = GetClaims()["sub"].(string)
	return
}
