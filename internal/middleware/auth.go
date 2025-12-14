package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"vasvault/pkg/utils"

	"github.com/gin-gonic/gin"
)

const UserCtxKey = "userID"

func BearerAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing Authorization header"))
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid Authorization header"))
			return
		}

		token := parts[1]
		claims, err := utils.ValidateAccessToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, claims.ID)
		next(w, r.WithContext(ctx))
	}
}

func GinBearerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
			return
		}

		token := parts[1]
		claims, err := utils.ValidateAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		fmt.Println("Token claims.ID:", claims.ID)
		c.Set("userID", claims.ID)
		c.Next()
	}
}
