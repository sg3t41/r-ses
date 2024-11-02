package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sg3t41/api/pkg/redis"
	"github.com/sg3t41/api/pkg/util/jwt"
)

func Get(c *gin.Context) {
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

	fmt.Println("***:")
	sessionID := claims.SessionID
	key := fmt.Sprintf("session:%s", sessionID)
	fmt.Println(key)

	//	isExists, err := redis.IsExists(c, key)
	//	if err != nil {
	//		fmt.Println("key:")
	//		fmt.Println(key)
	//		fmt.Println(err)
	//		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	//		return
	//	}

	githubAccessToken, err := redis.HGet(c, key, "github_access_token")
	if err != nil {
		fmt.Println(err)
		fmt.Println("key:")
		fmt.Println(key)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// GitHub APIにリクエストを送信
	client := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+githubAccessToken)
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request to GitHub API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "GitHub API request failed"})
		return
	}

	// GitHub APIからのレスポンスをデコード
	var repositories []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse GitHub API response"})
		return
	}

	// レポジトリ情報をJSONとして返す
	c.JSON(http.StatusOK, gin.H{"repositories": repositories})

}
