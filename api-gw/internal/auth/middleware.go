package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CtxUserKey = "auth.user"

func Middleware(cookieName string, store *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil || token == "" {
			c.Next()
			return
		}
		if sess, ok := store.Get(token); ok {
			c.Set(CtxUserKey, sess.User)
		}
		c.Next()
	}
}

func RequireAuth(cookieName string, store *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if sess, ok := store.Get(token); ok {
			c.Set(CtxUserKey, sess.User)
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}
