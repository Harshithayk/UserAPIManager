package service

import (
	"context"
	"errors"
	"fmt"
	"todo/pck/models"
	"todo/pck/utils"

	"github.com/redis/go-redis/v9"
)

func (s Service) UserSignup(ctx context.Context, userSigup models.UsersModel) (models.UserResponse, error) {

	ok1, _ := s.RepoUser.FindByEmail(ctx, userSigup.Email)
	if ok1 != (models.Users{}) {
		return models.UserResponse{}, errors.New("user is already exist")
	}

	ok := utils.ValidatePassword(userSigup.Password)
	if !ok {
		return models.UserResponse{}, errors.New("please provid the valid password")
	}
	user, err := utils.HashPassword(userSigup.Password)
	if err != nil {
		return models.UserResponse{}, errors.New("error when hashing the password")
	}

	userSig := models.Users{
		FirstName:   userSigup.FirstName,
		LastName:    userSigup.LastName,
		Email:       userSigup.Email,
		PhoneNumber: userSigup.PhoneNumber,
		Password:    user,
	}
	users, err := s.RepoUser.UserSigup(ctx, userSig)
	if err != nil {
		return models.UserResponse{}, errors.New("error when creating the user")
	}

	UserResponse := models.UserResponse{
		Id:      int(users.ID),
		Message: "user is created succssufully",
	}
	return UserResponse, nil
}

func (s Service) UserLogin(ctx context.Context, userLogin models.UserLogin) (models.LoginResponse, error) {

	userLog := models.Login{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	}
	paw, err := s.RepoUser.UserLogin(ctx, userLog)
	if err != nil {
		return models.LoginResponse{}, err
	}

	Token, err := s.Auth.GenerateJWT(userLogin.Email,paw.Role)
	if err != nil {
		return models.LoginResponse{}, err
	}

	err = utils.ComparePaswords(userLogin.Password, paw.Password)
	if err != nil {
		return models.LoginResponse{
			Message: "please enter the valid password ",
		}, err
	}
	
	err = s.rdb.AddTokenToCache(ctx, paw.Email, Token)
	if err != nil {
		return models.LoginResponse{}, err
	}
	cacheToken, err := s.rdb.GetTokenFromCache(ctx, paw.Email)
	fmt.Println("token from cache", cacheToken)
	if err == redis.Nil {
		cacheToken = Token
	} else if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Message: "logind in succussfly ",
		Tocken:  Token,
	}, nil
}

func (s Service) FetchUser(ctx context.Context) ([]models.FetchUser, error) {
	user, err := s.RepoUser.FetchUserrepo(ctx)
	if err != nil {
		return []models.FetchUser{}, err
	}
	return user, nil
}

func (s Service) FetchById(ctx context.Context, userId models.FetchByID) (models.FetchUser, error) {
	userdetails, err := s.RepoUser.FetchById(ctx, userId)
	if err != nil {
		return models.FetchUser{}, err
	}
	return userdetails, nil
}

func (s Service) UpdateUserById(ctx context.Context, UpdateId models.FetchUser) (models.UserResponse, error) {
	update, err := s.RepoUser.UpdateUserById(ctx, UpdateId)
	if err != nil {
		return models.UserResponse{}, err
	}
	return update, nil
}

func (s Service) DeleteById(ctx context.Context, FetchID int) (models.UserResponse, error) {
	resp, err := s.RepoUser.DeleteById(ctx, FetchID)
	if err != nil {
		return models.UserResponse{}, err
	}
	return resp, nil
}
