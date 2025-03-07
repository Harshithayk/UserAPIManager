package auth

const SingNature = "HiiIamHvingTheBadFrindTHATisNikiTha"

type Auth struct {
	SingNature string
}

type Authenticate interface {
	GenerateJWT(email string,relo string) (string, error)
	ValidateToken(token string) (Claims, error)
}

func NewAuth(SingNature string) (*Auth, error) {
	return &Auth{
		SingNature: SingNature,
	}, nil
}
