package auth

import "github.com/golang-jwt/jwt/v5"

const SingNature = "HiiIamHvingTheBadFrindTHATisNikiTha"

type Auth struct {
	SingNature string
}

type Authenticate interface {
	GenerateJWT(email string) (string, error)
	ValidateToken(token string) (jwt.RegisteredClaims, error)
}

func NewAuth(SingNature string) (*Auth, error) {
	return &Auth{
		SingNature: SingNature,
	}, nil
}
