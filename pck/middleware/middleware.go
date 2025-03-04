package middleware

import auth "todo/Auth"

type Middleware struct {
	Auth auth.Authenticate
}

func NewMiddleware(Auth auth.Authenticate) (*Middleware, error) {
	return &Middleware{
		Auth: Auth,
	}, nil
}
