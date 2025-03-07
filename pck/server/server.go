package server

import (
	"todo/pck/handlers"
	"todo/pck/middleware"

	"github.com/gin-gonic/gin"
)

func StartingServer(handlers *handlers.Handlers, m *middleware.Middleware) {
	router := gin.Default()
	router.POST("/signup", handlers.UserSignup)
	router.POST("/login", handlers.LoginPage)
	// Protected routes (Require authentication)
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(m.Authenticate())

	// Admin Routes (Only Admins Can Access)
	adminRoutes := protectedRoutes.Group("/admin")
	adminRoutes.Use(m.RoleAuthMiddleware("admin")) // ✅ Only admin can access
	{
		router.GET("/fetchUser",handlers.FetchAllUser)
		router.GET("/fetchById", handlers.FetchById)
		router.PUT("/updateById", handlers.UpdateUserById)
		router.DELETE("/DeleteById/:id", handlers.DeleteById)
	}

	// User Routes (Users & Admins Can Access)
	userRoutes := protectedRoutes.Group("/user")
	userRoutes.Use(m.RoleAuthMiddleware("user", "admin")) // ✅ Both users & admins can access
	{
		userRoutes.GET("/fetchById", handlers.FetchById)
	}
	router.Run(":8082")
}
