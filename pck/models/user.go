package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FirstName   string `gorm:"column:first_name;" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Email       string `gorm:"column:email; unique" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;" json:"phone_number"`
	Password    string `gorm:"column:password" json:"password"`
}

type UsersModel struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_Number" validate:"required"`
	Password    string `json:"password" validate:"min=8,max=15"`
}

type UserResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type Login struct {
	gorm.Model
	Email    string `gorm:"column:email; unique" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Tocken  string `json:"Token"`
}

type FetchUser struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_Number" validate:"required"`
}

type FetchByID struct {
	Id int `json:"id"`
}
