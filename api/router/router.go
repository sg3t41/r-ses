package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/middleware"
	"github.com/sg3t41/api/router/api/auth/github/callback"
	"github.com/sg3t41/api/router/api/auth/github/login"
	"github.com/sg3t41/api/router/api/v1/github/repositories/portfolio"
	repositories "github.com/sg3t41/api/router/api/v1/github/repositories/public"
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
		AllowHeaders:     []string{"Content-Type", "Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}
	r.Use(cors.New(corsConfig))

	r.GET("/api/auth/github/login", login.Get)
	r.GET("/api/auth/github/callback", callback.Get)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.GET("/github/repositories/public", repositories.Get)
		v1.POST("/github/repositories/portfolio", portfolio.Post)
	}
	return r
}
