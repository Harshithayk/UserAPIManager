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

func (m *Middleware) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
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

		next(c)

	}
}
