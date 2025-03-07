package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	email string
	role  string
	jwt.RegisteredClaims
}

func (a *Auth) GenerateJWT(email string, role string) (string, error) {
	claims := Claims{
		role:  role,
		email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    "LearnProject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.SingNature))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(token string) (Claims, error) {
	var c Claims
	tkn, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.SingNature), nil
	})
	if err != nil {
		return Claims{}, fmt.Errorf("error in parsing the token : %w", err)
	}

	if !tkn.Valid {
		return Claims{}, errors.New("token in not valid")
	}

	return c, nil

}
