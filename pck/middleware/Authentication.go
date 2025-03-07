package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ctxkey int

const key ctxkey = 1

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		authHeader := c.Request.Header.Get("Authorization")

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {

			err := errors.New("expected authorization header format: Bearer <token>")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err := m.Auth.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		ctx = context.WithValue(ctx, key, claims)

		req := c.Request.WithContext(ctx)
		c.Request = req

		c.Next()

	}
}

func (m *Middleware) RoleAuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the role from the Gin context
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied: no role found"})
			return
		}

		// Convert to string (Type Assertion)
		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role type"})
			return
		}

		// Check if the user role is in the allowed roles
		for _, allowedRole := range requiredRoles {
			if roleStr == allowedRole {
				c.Next() // User is authorized, proceed to the next handler
				return
			}
		}

		// If no match, deny access
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
	}
}
