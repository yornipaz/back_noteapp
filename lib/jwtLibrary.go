package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaimsCustom struct {
	Id      string
	Expired int64
	Secret  string
	Payload map[string]interface{}
}
type IJwtLibrary interface {
	Create(jwtClaimsCustom JwtClaimsCustom) (token string, err error)
	Parse(tokenString string) (*jwt.Token, error)
	ValidateExpiration(token *jwt.Token) bool
	ValidateToken(tokenString string) bool
	GetClaims(token *jwt.Token) jwt.MapClaims
}

type JwtLibrary struct {
}

// Parse implements IJwtLibrary.
func (*JwtLibrary) Parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

// ValidateExpiration implements IJwtLibrary.
func (*JwtLibrary) ValidateExpiration(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	expiration, ok := claims["exp"].(float64)
	if !ok {
		return false
	}

	return float64(time.Now().Unix()) <= expiration
}

// ValidateToken implements IJwtLibrary.
func (lib *JwtLibrary) ValidateToken(tokenString string) bool {
	token, err := lib.Parse(tokenString)
	if err != nil {
		return false
	}

	return lib.ValidateExpiration(token)
}

// create token wit the payload data
func (jwtLibrary *JwtLibrary) Create(jwtClaimsCustom JwtClaimsCustom) (token string, err error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     jwtClaimsCustom.Id,
		"exp":     time.Now().Add(time.Hour * time.Duration(jwtClaimsCustom.Expired)).Unix(),
		"payload": jwtClaimsCustom.Payload,
	})

	// Sign and get the complete encoded token as a string using the secret
	token, err = newToken.SignedString([]byte(jwtClaimsCustom.Secret))

	return
}
func (jwtLibrary *JwtLibrary) GetClaims(token *jwt.Token) jwt.MapClaims {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
func NewJwtLibrary() IJwtLibrary {
	return &JwtLibrary{}
}
