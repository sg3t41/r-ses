package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/pkg/util/jwt"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Authorizationヘッダーを取得
		tokenWithPrefix := c.GetHeader("Authorization")
		if tokenWithPrefix == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
		}

		token := strings.TrimPrefix(tokenWithPrefix, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}

		// JWTトークンの検証を行う
		claims, err := jwt.ParseToken(token)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}

		sessionID := claims.SessionID
		userID := claims.UserID
		c.Set("sessionID", sessionID)
		c.Set("userID", userID)

		c.Next()
	}
}
