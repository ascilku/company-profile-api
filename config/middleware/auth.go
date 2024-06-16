package middleware

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	GenerateToken(userID int) (string, error)
	ValidationToken(encodedToken string) (*jwt.Token, error)
}
type authMiddleware struct {
}

func NewAuthMiddleware() *authMiddleware {
	os.Getenv("SIGNED_KEY")
	return &authMiddleware{}
}

var SIGNED_KEY = []byte(os.Getenv("SIGNED_KEY"))

func (a *authMiddleware) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString(SIGNED_KEY)
	if err != nil {
		return signedString, err
	}
	return signedString, nil
}

func (a *authMiddleware) ValidationToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SIGNED_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
