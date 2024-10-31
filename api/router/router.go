package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors" // CORSパッケージのインポート
	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/pkg/redis"
	"github.com/sg3t41/api/pkg/util/jwt"
	"github.com/sg3t41/api/router/api/auth/github/callback"
	"github.com/sg3t41/api/router/api/auth/github/login"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS設定
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}

	r.Use(cors.New(corsConfig))

	r.GET("/", func(c *gin.Context) {
		redis.Set(c, "name", "john")
		redis.Get(c, "name")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/api/auth/github/login", login.Get)
	r.GET("/api/auth/github/callback", callback.Get)

	// GETメソッドでリポジトリを取得するエンドポイント
	r.GET("/api/v1/repositories/public", func(c *gin.Context) {

		// Authorizationヘッダーを取得
		tokenWithPrefix := c.GetHeader("Authorization")
		if tokenWithPrefix == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		token := strings.TrimPrefix(tokenWithPrefix, "Bearer ")

		// JWTトークンの検証を行う
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		fmt.Println("****:")
		fmt.Println(claims)

		repositories := []string{"Repo1", "Repo2", "Repo3"}

		c.JSON(http.StatusOK, gin.H{"repositories": repositories})
	})

	//	apiv1 := r.Group("/api/v1")

	return r
}
