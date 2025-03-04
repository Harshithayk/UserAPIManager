package handlers

import (
	"todo/pck/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Ser service.UserService
}

type UserHandler interface {
	Usersignup(ctx *gin.Context)
	LoginPage(ctx *gin.Context)
	FetchAllUser(ctx *gin.Context)
	UpdateUserById(ctx *gin.Context)
	DeleteById(ctx *gin.Context)
	FetchById(ctx *gin.Context)
}

func NewHandler(ser service.UserService) (*Handlers, error) {
	return &Handlers{
		Ser: ser,
	}, nil
}
