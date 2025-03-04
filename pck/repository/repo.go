package repository

import (
	"context"
	"todo/pck/models"

	"gorm.io/gorm"
)

type Repo struct {
	repo *gorm.DB
}

type UserRepository interface {
	UserSigup(ctx context.Context, user models.Users) (models.Users, error)
	UserLogin(ctx context.Context, userlog models.Login) (string, error)
	FetchUserrepo(ctx context.Context) ([]models.FetchUser, error)
	FetchById(ctx context.Context, id models.FetchByID) (models.FetchUser, error)
	UpdateUserById(ctx context.Context, update models.FetchUser) (models.UserResponse, error)
	DeleteById(ctx context.Context, delete int) (models.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (models.Users, error)
}

func Newrepository(db *gorm.DB) (*Repo, error) {
	return &Repo{
		repo: db,
	}, nil
}
