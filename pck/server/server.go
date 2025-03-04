package server

import (
	"todo/pck/handlers"
	"todo/pck/middleware"

	"github.com/gin-gonic/gin"
)

func StartingServer(handlers *handlers.Handlers, A *middleware.Middleware) {
	router := gin.Default()
	router.POST("/signup", handlers.UserSignup)
	router.POST("/login", handlers.LoginPage)
	router.GET("/fetchUser", A.Authenticate(handlers.FetchAllUser))
	router.GET("/fetchById", handlers.FetchById)
	router.PUT("/updateById", handlers.UpdateUserById)
	router.DELETE("/DeleteById/:id", handlers.DeleteById)
	router.Run(":8082")
}
