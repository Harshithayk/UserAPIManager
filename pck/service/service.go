package service

import (
	"context"
	auth "todo/Auth"
	"todo/pck/models"
	"todo/pck/repository"
)

type Service struct {
	RepoUser repository.UserRepository
	Auth     auth.Authenticate
}
type UserService interface {
	UserSignup(ctx context.Context, userSigup models.UsersModel) (models.UserResponse, error)
	UserLogin(ctx context.Context, userLogin models.UserLogin) (models.LoginResponse, error)
	FetchUser(ctx context.Context) ([]models.FetchUser, error)
	FetchById(ctx context.Context, FetchID models.FetchByID) (models.FetchUser, error)
	UpdateUserById(ctx context.Context, UpdateId models.FetchUser) (models.UserResponse, error)
	DeleteById(ctx context.Context, FetchID int) (models.UserResponse, error)
}

func NewService(RepoUser repository.UserRepository, auth auth.Authenticate) (*Service, error) {
	return &Service{
		RepoUser: RepoUser,
		Auth:     auth,
	}, nil
}
