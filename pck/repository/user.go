package repository

import (
	"context"
	"errors"
	"todo/pck/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) UserSigup(ctx context.Context, user models.Users) (models.Users, error) {
	users := r.repo.Create(&user)
	if users.Error != nil {
		return models.Users{}, errors.New("error when creating the user")
	}
	return user, nil
}

func (r *Repo) UserLogin(ctx context.Context, userlog models.Login) (user models.UserLogin, err error) {
	if err = r.repo.WithContext(ctx).Table("users").Select("*").Where("email= ?", userlog.Email).Find(&user).Error; err != nil {
		log.Error().Err(errors.New("CAN OUT FETCH THE DATA"))
		return models.UserLogin{}, err
	}
	return user, nil
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (models.Users, error) {
	var users models.Users
	if err := r.repo.WithContext(ctx).Table("users").Select("*").Where("email = ?", email).Find(&users).Error; err != nil {
		log.Error().Err(errors.New("CAN OUT FETCH THE DATA"))
		return models.Users{}, err
	}
	return users, nil
}

func (r *Repo) FetchUserrepo(ctx context.Context) ([]models.FetchUser, error) {
	var fetchUser []models.FetchUser
	if err := r.repo.WithContext(ctx).Table("users").Select("*").Find(&fetchUser).Error; err != nil {
		log.Error().Err(errors.New("CAN OUT FETCH THE DATA"))
		return []models.FetchUser{}, err
	}
	return fetchUser, nil
}

func (r *Repo) FetchById(ctx context.Context, id models.FetchByID) (models.FetchUser, error) {
	var fetchById models.FetchUser
	if err := r.repo.WithContext(ctx).Table("users").Select("*").Where("id=?", id.Id).Find(&fetchById).Error; err != nil {
		log.Error().Err(errors.New("CAN OUT FETCH THE DATA"))
		return models.FetchUser{}, err
	}
	return fetchById, nil
}

func (r *Repo) UpdateUserById(ctx context.Context, update models.FetchUser) (models.UserResponse, error) {
	//var updateid models.FetchUser

	if err := r.repo.WithContext(ctx).
		Table("users").
		Where("id = ?", update.Id).
		Updates(map[string]interface{}{
			"email":        update.Email,
			"first_name":   update.FirstName,
			"last_name":    update.LastName,
			"phone_number": update.PhoneNumber,
		}).Error; err != nil {

		log.Error().Err(err).Msg("FAILED TO UPDATE USER DATA")
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		Id:      update.Id,
		Message: "succussfully updated",
	}, nil
}

func (r *Repo) DeleteById(ctx context.Context, id int) (models.UserResponse, error) {
	if err := r.repo.WithContext(ctx).Table("users").
		Where("id = ?", id).
		Delete(&models.Users{}).Error; err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		Id:      id,
		Message: "deleted succussufully",
	}, nil
}
