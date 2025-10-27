package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CtxUserKey = "auth.user"

func Middleware(cookieName string, store SessionStore) gin.HandlerFunc { // <— interface
	return func(c *gin.Context) {
		if token, err := c.Cookie(cookieName); err == nil && token != "" {
			if user, ok := store.Get(token); ok {
				c.Set(CtxUserKey, user)
			}
		}
		c.Next()
	}
}

func RequireAuth(cookieName string, store SessionStore) gin.HandlerFunc { // <— interface
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if user, ok := store.Get(token); ok {
			c.Set(CtxUserKey, user)
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}
