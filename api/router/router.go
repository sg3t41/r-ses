package routers

import (
	"net/http"

	"github.com/gin-contrib/cors" // CORSパッケージのインポート
	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/pkg/redis"
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

	//	apiv1 := r.Group("/api/v1")

	return r
}
